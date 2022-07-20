d, err := os.Open("/Perkins") 
//fmt.Println(d.Readdirnames(-1)) 

y, err:=d.Readdirnames(-1) // this is the root 
fmt.Println(y) 

for i:=0; i<len(y); i++{ 
	if y[i]!=" "{
		 Folders:=y[i] 
		 temp,err:=os.Open("/Perkins"+Folders) //temp,err:=os.Open("/"Folders)
		 z,err:=temp.Readdirnames(-1) 
		 fmt.Println(y[i], " have the following subdirectories:") 
		 fmt.Println(z)for i:=0;i<len(Element2);i++{
			 if strings.Contains(Element2[i],"."){
				 test,err:=ioutil.Readfile("/Perkins"+Folders+"/"+z[(len(z)-1)]+"/"+Element2[i])
				  fmt.Println(" check the content of ", string(test))
			 }
			}
		}
	}

// ugh

package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func main() {
  filepath.Walk("/Perkins", func(path string, info os.FileInfo, err error) error {
    fmt.Println(path)
    return nil
  })
}


matched, err := filepath.Match("*.mp3", fi.Name())