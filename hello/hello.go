package hello

import (
    "log"
    "os"
    "hello_world/greetings"
    "hello_world/customLog"
)

func Run() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.

    // like this :
    // log.SetPrefix("greetings: ")
    // log.SetFlags(0)

    // or, custom like this :
    stdErr := log.New(os.Stderr, "", 0)
    stdOut := log.New(os.Stdout, "", 0)

    // Get a greeting message and print it.
    // with slice
    listNames := []string{
        "amar",
        "Udin",
        "cecep",
        "Budi",
    }
    messages, err1 := greetings.Hellos(listNames)
    // or without slice
    message, err2 := greetings.Hello("Amardito")
    // If an error was returned, print it to the console and
    // exit the program.
    if err1 != nil {
        //use log new standard (stdXxx) on first parameter customLog
        //second and last param is the module with exported function and the return
        customLog.Error(stdErr, "greetings.Hellos", err1)
        return
    }
    if err2 != nil {
        customLog.Error(stdErr, "greetings.Hello", err2)
        return
    }

    // If no error was returned, print the returned message
    // to the console.
    customLog.Success(stdOut, "greetings", messages)
    customLog.Success(stdOut, "greetings", message)
}
