open Format
open Advent

let load input = 
  Str.split (Str.regexp "\n") input
  |> List.map (fun line ->
    Str.split (Str.regexp ",") line
    |> List.map (fun part -> 
      Str.split (Str.regexp "-") part
      |> (fun part ->
          match part with
          | [x ; y] -> (int_of_string x, int_of_string y)
          | _ -> failwith "taco"
      )
    )
  )
;;

let surrounds = function
  | [(x, y) ; (x', y')] ->
    if (x >= x' && y <= y')
    || (x' >= x && y' <= y) 
    then 1 else 0
  | _ -> failwith "oops"
;;

let overlaps = function
  | [(x, y) ; (x', y')] ->
    if (x >= x' && x <= y')
      || (y >= x' && y <= y')
      || (x' >= x && x' <= y)
      || (y' >= x && y' <= y)
    then 1 else 0
  | _ -> failwith "oops"
;;


load d4sample |> List.map surrounds |> List.fold_left (+) 0 |> printf "%d\n";;
load d4input |> List.map surrounds |> List.fold_left (+) 0 |> printf "%d\n";;

load d4sample |> List.map overlaps |> List.fold_left (+) 0 |> printf "%d\n";; 
load d4input |> List.map overlaps |> List.fold_left (+) 0 |> printf "%d\n";;
