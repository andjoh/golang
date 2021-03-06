// +build !solution

// Leave an empty line above this comment.
package config
import ("strings" 
        "strconv"
       )


func (c *Configuration) parse(line string) (err error) {
  
    line=strings.TrimPrefix(line, "Number=")
    line=strings.TrimPrefix(line, "Name=")

    if num,err:=strconv.Atoi(line);err==nil{
      c.Number=num
    }else if len(line)>0{
      c.Name=line
    }
	return
}
