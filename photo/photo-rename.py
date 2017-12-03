#! /usr/bin/env python

"""
Rename files from <prefix><dont care>nnn.<ext> to <replace>-nnn.<ext>

Often the file names of scanned photos have the form like:
04110nnn.jpg
where 'nnn' is the picture number. I prefer to prefix them
<camera>-<film number>-nnn.jpg

This script does this.
"""

import sys
import os

def main():
    prefix = sys.argv[1]
    replace = sys.argv[2]

    files = os.listdir(".")

    for orig_name in files:
        if orig_name.startswith(prefix):
            name, _, ext = orig_name.rpartition('.')
            num = name[-3:]

            new_name = "%s-%s.%s" % (replace, num, ext)

            print "rename %s -> %s" % (orig_name, new_name)
            os.rename(orig_name, new_name)

if __name__ == '__main__':
    main()
