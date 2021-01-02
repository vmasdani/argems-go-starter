#!/usr/bin/env python3

import subprocess
import sys

if len(sys.argv) <= 1:
    print("Usage: ./release.py [base url]")
else:
    base_url = sys.argv[1]

    index_html_contents = """<html>
    <head>
    <meta charset="UTF-8">
    <title>Main</title>
    <script src="main.js"></script>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"> 
    </head>

    <body>
    <div id="myapp"></div>
    <script>
    var app = Elm.Main.init({{
        node: document.getElementById('myapp'),
        flags: '{}'
    }});
    </script>
    </body>
</html>""".format(base_url)

    steps = [
        ("mkdir -p dist/frontend", "."),
        ("go build", "."),
        ("cp argems-go-starter dist", "."),
        ("./build.sh", "./frontend"),
        ("cp dist/* ../dist/frontend", "./frontend"),
    ]

    for (cmd, cwd) in steps:
        subprocess.run(cmd, shell=True, cwd=cwd)

    f = open('./dist/frontend/index.html', 'w+')
    f.write(index_html_contents)
    f.close()

    print("\nRelease success! files are located in ./dist")
