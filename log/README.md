# Go Log

Log is very important for applications, it can illustrate the running status of the application in real time, and help to debug when the application has errors.
This example compares several log packages for Go and demonstrates how to use them.

There are 3 frequently used log packages:
- [log](https://github.com/golang/go/tree/master/src/log): The standard log package for Go.
   - Complex: Low
   - Advantages
      - Support prefix
      - Support code line number with short and long file model
      - Support rich flags for log format
   - Disadvantages
      - The long file model is the full path, but usually package path is enough
      - Only support 3 levels: Print, Fatal, Panic
- [glog](https://github.com/golang/glog): The leveled execution logs for Go from Google.
   - Complex: Middle
   - Advantages
      - Support 4 levels: Info, Warning, Error, Fatal
      - Support V log to customize level for Info
      - Support code line number
   - Disadvantages
      - Depend on `flag` package for log flags
      - Support code line number just with file name, without package path
- [logrus](https://github.com/sirupsen/logrus): The structured logger for Go.
   - Complex: High
   - Advantages
      - Support rich levels: Print, Debug, Info, Warning, Error, Fatal, Panic
      - Suport text and JSON format
      - Support fields
      - Support hooks for logging levels
   - Disadvantages
      - Not support code line number
      - Field structure is easy to be parsed by analysis tools, but hard to be read by human
