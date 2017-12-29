#!/usr/bin/env python3

import sys
import os
from os import path
from tempfile import mkstemp
from subprocess import check_output, check_call

GOPATH = os.getenv ('GOPATH')

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
			body, pre, #legend span {
				font-family: Menlo, monospace;
				font-weight: bold;
			}
			#topbar {
				background: black;
				position: fixed;
				top: 0; left: 0; right: 0;
				height: 42px;
				border-bottom: 1px solid rgb(80, 80, 80);
			}
			#content {
				margin-top: 50px;
			}
			#nav, #legend {
				float: left;
				margin-left: 10px;
			}
			#legend {
				margin-top: 12px;
			}
			#nav {
				margin-top: 10px;
			}
			#legend span {
				margin: 0 5px;
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
        <title>JCMS Tests Coverage</title>
	</head>
	<body>
        <div id="content">
            <table>
'''

HTML_TAIL = '''
            </table>
        </div>
    </body>
</html>
'''

COVDONE = '''
<tr>
    <td><a class="cov10" href="{}">{}</a></td>
    <td><span class="cov10">{}</span></td>
</tr>
'''
COVMISS = '''
<tr>
    <td><span class="cov0">{}</span>
    <td><span class="cov0">[no test files]</span></td>
</tr>
'''

def testcover (pkg):
    oldwd = os.getcwd ()
    os.chdir (path.join (GOPATH, 'src', pkg))
    dnfh = open (os.devnull, 'w')
    outfd, outfn = mkstemp (prefix = 'go-jcms.test.coverage')
    check_call ('go test -coverprofile coverage.out'.split (),
            stderr = dnfh, stdout = outfd)
    dnfh.close ()
    if path.isfile ('coverage.out'):
        check_call ('go tool cover -html coverage.out -o coverage.html'.split ())
        covdone (pkg, outfn)
    else:
        covmiss (pkg)
    fh = open (outfn, 'r')
    print (fh.read (), end = '')
    fh.close ()
    os.unlink (outfn)
    os.chdir (oldwd)

def covdone (pkg, outfn):
    covhtml = path.join (GOPATH, 'src', pkg, 'coverage.html')
    covinfo = ''
    fh = open (outfn, 'r')
    for line in [l.strip () for l in fh.readlines ()]:
        if line.startswith ('coverage: ') and line.endswith (' of statements'):
            covinfo = line
            break
    fh.close ()
    print (COVDONE.format (covhtml, pkg, covinfo), file = INDEX_FH)

def covmiss (pkg):
    print (COVMISS.format (pkg), file = INDEX_FH)

if __name__ == '__main__':
    global INDEX_FH
    INDEX_FH = open ('coverage.html', 'w')
    print (HTML_HEAD, file = INDEX_FH)
    if len (sys.argv) < 2:
        for pkg in check_output(['go', 'list', './...']).decode().splitlines():
            testcover (pkg)
    else:
        testcover (path.join ('github.com', 'jrmsdev', 'go-jcms', path.relpath (sys.argv[1])))
    print (HTML_TAIL, file = INDEX_FH)
    INDEX_FH.close ()
    sys.exit (0)
