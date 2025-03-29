#! /usr/bin/env python3

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
    reverse = False
    if len(sys.argv) > 3:
        reverse = True

    ofiles = os.listdir(".")

    files = [ f for f in ofiles if f.startswith(prefix) ]
    files.sort()
    
    idx = len(files)
    
    for orig_name in files:
        name, _, ext = orig_name.rpartition('.')
        num = name[-3:]

        if reverse:
            new_name = f'{replace}-{idx:03d}.{ext}'
        else:
            new_name = f'{replace}-{num}.{ext}'

        print(f'rename {orig_name} -> {new_name}')
        os.rename(orig_name, new_name)

        idx = idx - 1

if __name__ == '__main__':
    main()
