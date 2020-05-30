# await-or-not

Due to popular demand, we are releasing this library.

![Go!](./assets/mascot.png?raw=true "Go!")

## About

Powerful helpers to run a function async vs sync.

## Disclaimer

Guaranteed to work in Full Cyber Multi-INT Fusion envionments. Half Cyber unsupported.

## Typescript

## Javascript

## Go

### Usage

```
  import (
    . "github.com/apptness/await-or-not"
  )

  func Main() {
    // runs the lambda via asynchronously go routine
    //
    GoOrNot(true, func() {
      println("async!")
    })
    // runs the lambda synchronously
    AwaitOrNot(true, func() {
      println("sync!")
    })
  }
```
