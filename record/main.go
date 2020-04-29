package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

func main() {
	in := bytes.Buffer{}
	outInfo := bytes.Buffer{}
	cmd := exec.Command("cmd")
	//cmd := exec.Command("ping", "www.163.com")

	cmd.Stdout = &outInfo
	cmd.Stdin = &in
	//in.WriteString("echo hello world > test.txt\n")
	// go func() {
	// 	//in.WriteString("ffmpeg -f dshow -i  audio='麦克风 (Realtek(R) Audio)' -ar 16000 -ac 1 -t 00:00:05 test.wav\n")
	// 	//in.WriteString("echo hello world > test.txt\n")
	// }()

	in.WriteString("ffmpeg -f dshow -i  audio=@device_cm_{33D9A762-90C8-11D0-BD43-00A0C911CE86}\\wave_{34BE0A01-C298-4CAA-88F7-87EFA5B666A6} -ar 16000 -ac 1 -t 00:00:05 test.wav -y \n")
	//in.WriteString("ping www.baidu.com\n")
	//in.WriteString("echo hello world\n")
	// if err := cmd.Run(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = cmd.Wait(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(cmd.ProcessState.Pid())
		fmt.Println(cmd.ProcessState.ExitCode)
		fmt.Println(cmd.ProcessState.Sys().(syscall.WaitStatus).ExitCode)
		fmt.Println(outInfo.String())
	}

	//time.Sleep(5 * time.Second)
}

//转码

//func main() {
//	cmdArguments := []string{"-i", "divx.avi", "-c:v", "libx264",
//		"-crf", "20", "-c:a", "aac", "-strict", "-2", "video1-fix.ts"}
//
//	cmd := exec.Command("ffmpeg", cmdArguments...)
//
//	var out bytes.Buffer
//	cmd.Stdout = &out
//	err := cmd.Run()
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Printf("command output: %q", out.String())
//}

/*
func main() {
	// cmdArguments := []string{"-i", "divx.avi", "-c:v", "libx264",
	// 	"-crf", "20", "-c:a", "aac", "-strict", "-2", "video1-fix.ts", "-y"}

	cmdArguments := []string{"-f", "dshow", "-i", "audio=@device_cm_{33D9A762-90C8-11D0-BD43-00A0C911CE86}\\wave_{34BE0A01-C298-4CAA-88F7-87EFA5B666A6}",
		"-ar", "16000", "-ac", "1", "-t", "00:00:05", "test.wav", "-y"}
	//cmdArguments := []string{"-i", "divx.avi"}
	//cmd := exec.Command("ffmpeg", "-i", "divx.avi")

	//cmdArguments := []string{"-list_devices", "true", "-f", "dshow", "-i", "dummy"}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	w := bytes.NewBuffer(nil)

	out := bytes.Buffer{}
	cmd.Stdout = &out
	cmd.Stderr = w
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Printf("command output: %q", out.String())
	fmt.Printf("Stderr: %s\n", string(w.Bytes()))
}
*/
