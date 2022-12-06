open Format
open Advent

module Counts = Map.Make(Char);;

let char_pos c =
  let n = Char.code c in
  if n < 97 then n - 64 + 26 else n - 96

let dupes chars = 
  (* could use Hashtbl.lenght instead of n *)
  let rec aux acc n = function
    | [] -> acc
    | (char :: t) when n > 0 -> aux (Counts.add char 1 acc) (n-1) t
    | (char :: t) ->
      let counts = match (Counts.find_opt char acc) with
      | None -> acc
      | Some n -> Counts.add char (n+1) acc
      in aux counts n t
  in
  let n = List.length chars / 2
  in
  let counts = aux (Counts.empty) n chars
  in
  Counts.fold (fun k v acc->
    if v > 1 then (
      (* printf "%c %d %d\n" k v (char_pos k); *)
      acc + (char_pos k)
    ) else
      acc
  ) counts 0
;;

let load input = 
  Str.split (Str.regexp "\n") input
    |> List.map (fun line ->
    Str.split (Str.regexp "") line
      |> List.map (fun s -> 
          (* printf "%s\n" s; *)
          String.get s 0)
  )
;;

(* chunks [1;2;3;4;5;6];; *)
let chunks list =
  let rec aux acc = function
  | [] -> acc
  | a::b::c::tail -> aux (acc @ [[a;b;c]]) tail
  | _ -> failwith "nope"
  in
  aux [] list

 module SS = Set.Make(String);;
let common = function
  | [a;b;c] -> 
    let a = SS.of_list (Str.split (Str.regexp "") a) in
    let b = SS.of_list (Str.split (Str.regexp "") b) in
    let c = SS.of_list (Str.split (Str.regexp "") c) in
    SS.inter a b |> SS.inter c |> SS.elements (* |> List.iter (printf "%s\n") *)
  | _ -> failwith "nope"




(* dupes (load d3sample) |> printf "%d\n" *)
let values = load d3sample |> List.map dupes;;
(* List.iter (printf "%d\n") values;; *)
values |> List.fold_left (+) 0 |> printf "%d\n";;

let values = load d3input |> List.map dupes;;
(* List.iter (printf "%d\n") values;; *)
values |> List.fold_left (+) 0 |> printf "%d\n";;


Str.split (Str.regexp "\n") d3sample |> chunks |> List.map common |> List.map (fun x -> String.get (List.hd x) 0) |> List.map char_pos |> List.fold_left (+) 0 |> printf "%d\n";;
Str.split (Str.regexp "\n") d3input |> chunks |> List.map common |> List.map (fun x -> String.get (List.hd x) 0) |> List.map char_pos |> List.fold_left (+) 0 |> printf "%d\n"
(* Str.split (Str.regexp "\n") d3input |> chunks |> List.map common |> List.iter (fun x -> printf "%s" (List.hd x)) *)

