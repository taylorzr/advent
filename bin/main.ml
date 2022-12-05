open Format
open Advent

module SI = Set.Make(Int);;

let range x y = 
  let rec aux acc start stop =
    if start == (stop + 1) then
      acc
    else 
      aux (acc @ [start]) (start+1) stop
  in aux [] x y
;;


let load input = 
  Str.split (Str.regexp "\n") input
  |> List.map (fun line ->
    Str.split (Str.regexp ",") line
    |> List.map (fun part -> 
      Str.split (Str.regexp "-") part
      |> (fun part ->
          match part with
          | [x ; y] -> range (int_of_string x) (int_of_string y)
          | _ -> failwith "taco"
      )
    )
  )
;;

let contains = function
  | [ x ; y ] -> (
    let shorter = min (List.length x) (List.length y) in
    let common = SI.inter (SI.of_list x) (SI.of_list y) in
    if (SI.cardinal common) == shorter then 1 else 0
  )
  | _ -> failwith "oops"
;;

let overlaps = function
  | [ x ; y ] -> (
    let common = SI.inter (SI.of_list x) (SI.of_list y) in
    if (SI.elements common |> List.length) > 0 then 1 else 0
  )
  | _ -> failwith "oops"
;;

load d4sample |> List.map contains |> List.fold_left (+) 0 |> printf "%d\n";;
load d4input |> List.map contains |> List.fold_left (+) 0 |> printf "%d\n";;

load d4sample |> List.map overlaps |> List.fold_left (+) 0 |> printf "%d\n";;
load d4input |> List.map overlaps |> List.fold_left (+) 0 |> printf "%d\n";;


