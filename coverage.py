#!/usr/bin/env python3

import sys
import os
from os import path
from time import asctime
from tempfile import mkstemp
from subprocess import check_output, check_call, CalledProcessError

GOPATH = os.getenv ('GOPATH')
DOCROOT = path.abspath(path.join ('.', 'htmlcov'))
INDEX_FN = path.join(DOCROOT, 'index.html')

HTML_HEAD = '''
<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <style>
            body {
                background: black;
                color: rgb(80, 80, 80);
            }
            body, pre, span {
                font-family: Menlo, monospace;
                font-weight: bold;
            }
            .cov0 { color: rgb(192, 0, 0) }
            .cov1 { color: rgb(128, 128, 128) }
            .cov2 { color: rgb(116, 140, 131) }
            .cov3 { color: rgb(104, 152, 134) }
            .cov4 { color: rgb(92, 164, 137) }
            .cov5 { color: rgb(80, 176, 140) }
            .cov6 { color: rgb(68, 188, 143) }
            .cov7 { color: rgb(56, 200, 146) }
            .cov8 { color: rgb(44, 212, 149) }
            .cov9 { color: rgb(32, 224, 152) }
            .cov10 { color: rgb(20, 236, 155) }
        </style>
        <title>go-jcms tests overage</title>
    </head>
    <body>
        <div id="content">
            <table>
'''

HTML_TAIL = '''
            </table>
        </div>
        <br>
        <footer>
            <small>{}</small>
        </footer>
    </body>
</html>
'''

COVMISS = '''
<tr>
    <td>{}</td>
    <td>[no test files]</td>
</tr>
'''

COVERR = '''
<tr>
    <td><span class="cov0">{}</span>
    <td><span class="cov0">FAIL</span></td>
</tr>
'''

def testcover (pkg):
    oldwd = os.getcwd ()
    os.chdir (path.join (GOPATH, 'src', pkg))
    errfd, errfn = mkstemp (prefix = 'go-jcms.test.coverage.err')
    outfd, outfn = mkstemp (prefix = 'go-jcms.test.coverage.out')
    try:
        check_call ('go test -coverprofile coverage.out'.split (),
                stderr = errfd, stdout = outfd)
    except CalledProcessError:
        fh = open (errfn, 'r')
        print (fh.read (), end = '')
        fh.close ()
        os.unlink (errfn)
        fh = open (outfn, 'r')
        print (fh.read (), end = '')
        fh.close ()
        os.unlink (outfn)
        coverr (pkg)
        return
    if path.isfile ('coverage.out'):
        covdocroot = path.join(DOCROOT, pkg)
        covfilename = path.join(covdocroot, 'coverage.html')
        os.makedirs(covdocroot, exist_ok = True)
        cmd = 'go tool cover -html coverage.out -o {}'.format(covfilename)
        check_call (cmd.split ())
        covdone (pkg, outfn)
    else:
        covmiss (pkg)
    fh = open (outfn, 'r')
    print (fh.read (), end = '')
    fh.close ()
    os.unlink (outfn)
    os.unlink (errfn)
    os.chdir (oldwd)

COVDONE = '''
<tr>
    <td><a class="cov{covlevel}" href="{href}">{pkg}</a></td>
    <td><span class="cov{covlevel}">{covinfo}</span></td>
</tr>
'''

def covdone (pkg, outfn):
    href = path.join (DOCROOT, pkg, 'coverage.html')
    covinfo = ''
    fh = open (outfn, 'r')
    for line in [l.strip () for l in fh.readlines ()]:
        if line.startswith ('coverage: ') and line.endswith (' of statements'):
            covinfo = line
            break
    fh.close ()
    if covinfo == '':
        covmiss (pkg)
    else:
        covp = int(covinfo.strip().split()[1].split('.')[0])
        covlevel = int(covp / 10)
        print (COVDONE.format (href = href, pkg = pkg, covinfo = covinfo,
               covlevel = covlevel), file = INDEX_FH)

def covmiss (pkg):
    print (COVMISS.format (pkg), file = INDEX_FH)

def coverr (pkg):
    print (COVERR.format (pkg), file = INDEX_FH)

if __name__ == '__main__':
    global INDEX_FH
    os.makedirs (DOCROOT, exist_ok = True)
    INDEX_FH = open (INDEX_FN, 'w')
    print (HTML_HEAD, file = INDEX_FH)
    gopatt = './lib/...'
    if len (sys.argv) == 2:
        gopatt = path.join ('github.com', 'jrmsdev', 'go-jcms', sys.argv[1])
    for pkg in check_output(['go', 'list', gopatt]).decode().splitlines():
        testcover (pkg)
    now = asctime ()
    print (HTML_TAIL.format (now), file = INDEX_FH)
    INDEX_FH.close ()
    print (INDEX_FN)
    sys.exit (0)
