const { spawn } = require("child_process")
const os = require("os")
const path = require("path")

const env = { ...process.env }

const homeDir = os.homedir()
env.GOPATH = path.join(homeDir, "go")
env.GOMODCACHE = path.join(env.GOPATH, "pkg", "mod")
env.GOCACHE = path.join(env.GOPATH, "go-build")
const cmd = spawn("go", ["run", "main.go"], {
  stdio: "inherit",
  env,
})

cmd.on("exit", code => {
  process.exit(code)
})
