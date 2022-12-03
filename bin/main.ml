open Format
open Advent

let char_pos c =
  let n = Char.code c in
  if n < 97 then n - 64 + 26 else n - 96

let dupes string = 
  (* could use Hashtbl.lenght instead of n *)
  let rec aux acc n = function
    | [] -> acc
    | (char :: t) -> (if n > 0 then
      if Hashtbl.mem acc char then () else Hashtbl.add acc char 1
    else
      if Hashtbl.mem acc char then Hashtbl.replace acc char 2
    ); aux acc (n-1) t
  in
  let n = List.length string / 2
  in
  let counts = aux (Hashtbl.create n) n string
  in
  Hashtbl.fold (fun k v acc->
    if v > 1 then (
      printf "%c %d %d\n" k v (char_pos k);
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

(* dupes (load d3sample) |> printf "%d\n" *)
let values = load d3sample |> List.map dupes;;
List.iter (printf "%d\n") values;;
values |> List.fold_left (+) 0 |> printf "%d\n";;

let values = load d3input |> List.map dupes;;
(* List.iter (printf "%d\n") values;; *)
values |> List.fold_left (+) 0 |> printf "%d\n";;
