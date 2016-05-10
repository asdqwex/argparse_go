package argparse

import (
    "fmt"
    "strings"
)

type Argument struct {
    takes_input string
    longname  string
    required bool
    conflicts []string
    dependencies []string
    only_once bool
    reads_from_stdin bool
    hungry bool
}

// Example arg_def_map
//  my_apps_args = {
//    a = {
//      takes_input = 'string' # type of data this function needs
//      longname = 'function1' # if called with -- what is it called
//      required = 0           # does this option always need to be defined
//      conflicts = ['b', 'd'] # other options this function cannot be run with
//      dependencies = ['a']   # other options this function requires
//      only_once = 1          # this function should only be called once
//      reads_from_stdin: 0    # this function read from stdin
//      hungry = 1             # this function takes the rest of the input data
//    },
//    b = {
//      ...
//      ...
//    },
//    c = {
//      ...
//      ...
//    }
// }

func Parse(argument_array []string, arg_def_map map[string]string)) (arg_out_map map[string]string)  {
  // initialize arg_out_map
  arg_out_map = make(map[string]string)
  // initialize bool to keep track of wethere the next item in argument_array is an argument  or an inout value
  next_is_data_pass := 0
  // remove file name from argument_array
  argument_array = append(argument_array[:0], argument_string[1:]...)
  // loop over each item in argument_array
  for i := range argument_array {
    // if this item has been marked as a value skip
    if next_is_data_pass == 1 {
      // reset trackin bool
      next_is_data_pass = 0
      // skip
      continue
    }
    // if the item is only - take input from STDIN
    if argument_array[i] == "-" {
      // read from stdin to data value of the last flag? or maybe the function has to do the reading, tbd

      // assume done and break loop
      break
    }
    // if the item is only -- everything else is data
    if argument_array[i] == "--" {
      // stringify the rest of the argument_array and set as data for last flag? or maybe any flag marked to take the rest of the data?

      // assume done and break loop
      break
    }
    // if item is prefixed with -
    if strings.HasPrefix(argument_array[i], "-") {
      // remove - from argument
      argument_array[i] = argument_array[i][1:]
      // if argument contains =
      if strings.Contains(argument_array[i], "=") {
        // split on =
        p := strings.Split(argument_array[i], "=")
        // set key and value
        arg_out_map[p[0]] = p[1]
        // NEXT
        continue
      }
      // get index for next item in argument_array
      i2 := i + 1
      //set value to next item in argument_array if no other cases have matched
      next := argument_array[i2]
      // // set key and value
      arg_out_map[argument_array[i]] = next
      // do not treat the next item in argument_array as an argument
      next_is_data_pass = 1
    // catch all case
    } else {
      // set item as key with value of "action"
      arg_out_map[argument_array[i]] = "action"
    }
  }
  // initialize the action counter
  action_counter := 0
  // loop over each value in arg_out_map
  for _ , v := range arg_out_map {
    // check if value is set to action
    if v == "action" {
      // increment action counter if value is "action"
      action_counter++
    }
  }
  // see if the action_counter is greater that 1
  if action_counter > 1 {
    // if it is panic
    panic(fmt.Sprintf("%v", "there were more than on actions called"))
  }
  // pass the formatted arg_out_map to the program that called it
  return arg_out_map
}
