package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/pelletier/go-toml"
)

func panik(error error) {
	if error != nil {
		panic(error)
	}
}

func main() {
	rcv, err := os.ReadFile("./lccli.toml")
	panik(err)

	conf, err := toml.Load(string(rcv))
	panik(err)

	if len(os.Args) < 2 {
		fmt.Println("Usage: lccli <version>")
		os.Exit(1)
	}
	version := os.Args[1]
	switch version {
	case "1.7":
	case "1.8":
	case "1.12":
	case "1.16":
	case "1.17":
	case "1.18":
		break
	default:
		fmt.Println("Version not supported. Please use 1.7, 1.8, 1.12, 1.16, 1.17 or 1.18")
	}
	asset := version
	if asset == "1.7" {
		asset = "1.7.10"
	}

	h, err := os.UserHomeDir()
	panik(err)
	lunar := h + "/.lunarclient"
	game := conf.Get("Game").(*toml.Tree).Get(version).(*toml.Tree)
	dir := game.Get("dir").(string)
	java := game.Get("java").(string)
	cosmetics := game.Get("cosmetics").(bool)
	csmtcs := ""
	if dir == "" {
		fmt.Println("Directory not found for version " + version)
	}
	if java == "" {
		fmt.Println("Java not found for version " + version)
	}
	if cosmetics {
		csmtcs = lunar + "/textures"
	}
	if java == "default" {
		j := "java"
		if runtime.GOOS == "windows" {
			j += "w.exe"
		}
		path := fmt.Sprintf("%s/jre/%s/zulu*/bin/%s", lunar, version, j)
		f, err := filepath.Glob(path)
		panik(err)
		java = f[0]
	}
	_, err = os.Stat(java)
	panik(err)

	jvm := conf.Get("JVM").(*toml.Tree)
	jargs := jvm.Get("args").(string)
	if jargs == "" {
		fmt.Println("Arguments not found for JVM")
	}

	off := lunar + "/offline/" + version
	n := off + "/natives"
	libs := "--add-modules jdk.naming.dns --add-exports jdk.naming.dns/com.sun.jndi.dns=java.naming -Dlog4j2.formatMsgNoLookups=true -Djna.boot.library.path=" + n + " --add-opens java.base/java.io=ALL-UNNAMED"
	launch := "-Djava.library.path=" + n + " -XX:+DisableAttachMechanism -cp "
	jars := []string{"lunar-assets-prod-1-optifine.jar", "lunar-assets-prod-2-optifine.jar", "lunar-assets-prod-3-optifine.jar", "lunar-libs.jar", "lunar-prod-optifine.jar", "OptiFine.jar"}
	for _, jar := range jars {
		launch += fmt.Sprintf("%s/%s:", off, jar)
	}
	launch += fmt.Sprintf("%s/vpatcher-prod.jar com.moonsworth.lunar.patcher.LunarMain --version %s --accessToken 0 --assetIndex %s --userProperties {} --gameDir %s --width 854 --height 480 --texturesDir %s --assetsDir %s/assets", off, version, asset, dir, csmtcs, dir)
	launch = fmt.Sprintf("%s %s %s %s", java, libs, jargs, launch)
	fmt.Println(launch)
	shell := "bash"
	c := "-c"
	if runtime.GOOS == "windows" {
		shell = "cmd"
		c = "/C"
	}
	cmd := exec.Command(shell, c, launch)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())
}
