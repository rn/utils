#! /usr/bin/env python

"""Set various EXIF tags for scanned photos/negatives using exiftool"""

import getopt
import subprocess
import sys

# Disable bad whitespace warning
# pylint: disable=C0326

# Camera 'Make' and 'Model' for EXIF data
CAMERAS = {
    1: ["Leica", "Leica M7"],
    2: ["Zeiss", "Zeiss Ikon ZM"],
    3: ["Yashica", "Yashica-Mat"],
    4: ["Rollei", "Rollei 35S"],
    5: ["Fuji", "Fuji GW670 III 6x7 Professional"],
    6: ["Minolta", "Minolta XG-1"],
    7: ["Olympus", "Olympus-35 SP"],
    }

# Lens name and focal length
LENSES = {
    1: ["Zeiss Biogon 2/35 ZM", "35"],
    2: ["Leica Summicron-M 1:2/50", "50"],
    3: ["Leitz Tele-Elmarit-M 1:2.8/90", "90"],
    4: ["Zuiko 42mm f/1.7", "42"],
    5: ["Yashinon 1:3.5 80mm", "80"],
    6: ["Rollei-HFT Sonnar 2.8/40 ", "40"],
    7: ["Fujinon 1:3.5 90mm EBC", "90"],
    8: ["Minolta MD Rokkor 50mm 1:1.7", "50"],
    9: ["Minolta MD Zoom 28-70mm f/3.5-4.8", "unknown"],
    10: ["Vivitar MC Macro Focusing Zoom 70-210mm 1:4.4-5.6", "unknown"],
    }

# Film "ISO" and Film description and Negative inscription
FILMS = [
    ["400",  "Kodak TRI-X 400", ""],
    ["400",  "Kodak T-MAX 400", "TMY 5053"],
    ["100",  "Kodak T-MAX 100", ""],
    ["100",  "Fujifilm 100 Acros", "ACR-36"],
    ["400",  "Fujifilm Neopan 400", "400-PR"],
    ["1600", "Fujifilm Neopan 1600", "1600-PR"],
    ["400",  "Kodak T400 CN", "T400CN"],
    ["400",  "Illford HP5 Plus", ""],
    ["125",  "Illford HP4", ""],
    ["400",  "Illford XP2 400", ""],
    ["400",  "Illford 400 Delta Professional", ""],
    ["50",   "Fujichrome Velvia 50", "RVP-50"],
    ["100",  "Fujichrome Velvia 100", "RVP-100"],
    ["100",  "Fujichrome Velvia 100F", ""],
    ["100",  "Fujichrome Sensia 100", "RD-104"],
    ["200",  "Fujichrome Sensia 200", "RM-905"],
    ["100",  "Fujicolor Superia 100", "CN, S-100"],
    ["160",  "Fujicolor Pro 160 NS", "PN, 160NS"],
    ["200",  "Fujicolor Superia 200 CA", "CA-3, G-200"],
    ["400",  "Fujicolor Press/Superia X-TRA 400", "CH,S-400"],
    ["800",  "Fujicolor Press/Superia X-TRA 800", "CZ,G-800"],
    ["50",   "ADOX CHS 50", "CHS50"],
    ["100",  "ADOX CHS 100", ""],
    ["200",  "Kodak EktaChrome 200", ""],
    ["100",  "Kodak EktaChrome 100", "EB 5045"],
    ["100",  "Kodak Ektar 100-2", "CX 5301"],
    ["125",  "Kodak Ektar 125-1", ""],
    ["100",  "Kodak Gold 100-2", ""],
    ["200",  "Kodak Gold 200-2", "GB 6096 or GB 7304"],
    ["160",  "Kodak Portra 160VC", "160-VC2"],
    ["100",  "Kodak Color II 100", "Kodak Safety 5053"],
    ["40",   "Agfacolor CN 17", ""],
    ["80",   "Agfacolor Special CNS", ""],
    ["80",   "Agfacolor Special CNS2", ""],
    ["200",  "Agfa XRG 200", ""],
    ["125",  "Agfa Optima 125", ""],
    ["50",   "Agfa Ultra 50", ""],
    ["40",   "Agfa Leverkusen Isopan F", "AGFA L IF"],
    ["100",   "Agfa Isopan SS", "AGFA ISS"],
    ]

EXIFTOOL = "exiftool -overwrite_original"
et_opt = ""

def print_cameras():
    """Print a list of known cameras"""
    ks = CAMERAS.keys()
    ks.sort()
    print "Supported Camera Models:"
    print "Index Manufacturer Model"
    for k in ks:
        print "%5d %-12s %s" % (k, CAMERAS[k][0], CAMERAS[k][1])

def print_lenses():
    """Print a list of known lenses"""
    ls = LENSES.keys()
    ls.sort()
    print "Supported Lens Models:"
    print "Index Lens"
    for l in ls:
        print "%5d %s" % (l, LENSES[l][0])

def print_films():
    """Print a list of known films"""
    print "Supported Films:"
    print "Index  ISO Description"
    for f in FILMS:
        print "%5d %-4s %-30s %s" % (FILMS.index(f) + 1, f[0], f[1], f[2])

def usage():
    """Print usage"""
    print "scan-exif.py [-c <n>|h] [-l <n>|h] [-f <n>|h] <files>"
    print "Munge exif data from scanned negatives"
    print "-c <n>: Set Camera Model. h for list of cameras"
    print "-l <n>: Set Lens. h for list of lenses"
    print "-f <n>: Set Films and ISO. h for list of films"
    print "-i <n>: Override Film ISO"
    print "-d <n>: Date format: YYYY:MM:DD"
    print "-t <n>: Time format: HH:MM (increment MM for multiple files)"
    print "-a <n>: Aperture"

if __name__ == "__main__":
    arglist = "hc:l:f:i:d:t:a:"
    camera = 0
    lens = 0
    film = 0
    f_iso = ""
    iso = 0
    date = None
    hour = 12
    minute = 0
    aperture = None

    opts, files = getopt.getopt(sys.argv[1:], arglist, [])
    try:
        for opt in opts:
            if opt[0] == '-h':
                usage()
                sys.exit()
            if opt[0] == '-c':
                if opt[1] == 'h':
                    print_cameras()
                    sys.exit()
                else:
                    camera = int(opt[1])
            if opt[0] == '-l':
                if opt[1] == 'h':
                    print_lenses()
                    sys.exit(0)
                else:
                    lens = int(opt[1])
            if opt[0] == '-f':
                if opt[1] == 'h':
                    print_films()
                    sys.exit(0)
                else:
                    film = int(opt[1])
            if opt[0] == '-i':
                iso = int(opt[1])
            if opt[0] == '-d':
                date = opt[1]
            if opt[0] == '-t':
                h_s, _, m_s = opt[1].partition(':')
                hour = int(h_s)
                minute = int(m_s)
            if opt[0] == '-a':
                aperture = float(opt[1])
    except:
        raise

    if not camera == 0:
        c_man = CAMERAS[camera][0]
        c_mod = CAMERAS[camera][1]
        et_opt += ' -Make=%s -Model="%s"' % (c_man, c_mod)

    if not lens == 0:
        l_desc = LENSES[lens][0]
        l_foc = LENSES[lens][1]
        et_opt += ' -Lens="%s"' % (l_desc)
        et_opt += ' -FocalLength="%s"' % (l_foc)

    if not film == 0:
        f_iso = FILMS[film - 1][0]
        f_desc = FILMS[film - 1][1]
        et_opt += ' -HierarchicalSubject+="Film|%s"' % f_desc

    if not iso == 0:
        et_opt += " -ISO=%s" % iso
    elif not f_iso == "":
        et_opt += " -ISO=%s" % f_iso

    if aperture:
        et_opt += ' -FNumber=%.1f' % aperture


    for infile in files:

        et_opt_cur = et_opt
        if date:
            et_opt_cur += ' -DateTimeOriginal="%s %02d:%02d:00"' % \
                          (date, hour, minute)
            et_opt_cur += ' -CreateDate="%s %02d:%02d:00"' % \
                          (date, hour, minute)
            minute += 1
            if minute >= 60:
                hour += 1
                minute = 0

        cmd = "%s%s '%s'" % (EXIFTOOL, et_opt_cur, infile)
        print cmd
        subprocess.call(cmd, shell=True)
