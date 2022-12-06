open Format
open Advent

(* specify # stacks when loading input *)

let parse_commands input =
  (* let template = Str.regexp "move (\d) from (\d) to (\d)" in *)
  let tmpl = Str.regexp {|move \([0-9]\) from \([0-9]\) to \([0-9]\)|} in
  Str.split (Str.regexp "\n") input
  |> List.map (fun line ->
    if Str.string_match tmpl line 0 then
      let a, b, c = Str.matched_group 1 line, Str.matched_group 2 line, Str.matched_group 3 line in
      let () = printf "%s %s %s\n" a b c in
      ((int_of_string a), (int_of_string b), (int_of_string c))
    else
      failwith (sprintf "no match for '%s'" line)
  )
;;

let rec firstn n xs = match xs with
| [] -> failwith "firstn"
| x::xs ->
    if n=1 then [x] else x::firstn (n-1) xs;;

let rec dropn n xs = match xs with
| [] -> failwith "lastn"
| _::xs -> if n=1 then xs else dropn (n-1) xs
;;

 let rec transpose list = match list with
| []             -> []
| []   :: xss    -> transpose xss
| (x::xs) :: xss ->
    (x :: List.map List.hd xss) :: transpose (xs :: List.map List.tl xss)


(* let parse_stack input = *)
(*   let tmpl = Str.regexp {|.\([ A-Z]\). .\([ A-Z]\). .\([ A-Z]\).|} in *)
(*   Str.split (Str.regexp "\n") input *)
(*   |> List.fold_left (fun line -> *)
(*     if Str.string_match tmpl line 0 then *)
(*       let a, b, c = Str.matched_group 1 line, Str.matched_group 2 line, Str.matched_group 3 line in *)
(*       let () = printf "%s %s %s\n" a b c in *)
(*       [ List.nth stacks 0 @ [a] ; List.nth stacks 1 @ [b] ; List.nth stacks 2 @ [c] ] *)
(*     else *)
(*       failwith (sprintf "no match for '%s'" line) *)
(*   ) [[];[];[]] *)
(* ;; *)

let parse_stack input = 
  let positions = [1 ; 5 ; 9] in (* TODO: calc positions *)
  Str.split (Str.regexp "\n") input (* TODO: string column nums row *)
  |> List.map (fun line ->
    List.map (fun pos -> String.get line pos) positions
  )
  |> transpose
  |> List.map (List.filter (fun c -> c != ' '))
  (* |> List.iter (fun stack -> *)
  (*     List.iter (printf "%c") stack; *)
  (*     printf "\n" *)
  (* ) *)
;;



let process command stacks =
  let (n, t, g) = command in
  let taking_i, giving_i = (t - 1, g - 1) in
  let moving = List.nth stacks taking_i |> firstn n in
  List.mapi (fun i stack ->
    if i = giving_i then
      moving @ stack
    else if i = taking_i then
      dropn ((List.length stack) - n - 1) stack
    else
      stack
  ) stacks
;;

let rec process_all stacks = function
  | [] -> stacks
  | command :: tail -> process_all (process command stacks) tail
;;

let load input _ =
  Str.split (Str.regexp "\n\n") input
  |> (fun parts ->
      match parts with
      | [stack ; commands] -> (
        (parse_stack stack),
        (parse_commands commands)
      )
      | _ -> failwith "oops"
    )
;;

(* let print_stack stack = *)

let (stacks, commands) = load d5sample 3 in
(* List.fold_left (fun (n, t , g) -> process n t g stacks) stacks commands *)
process_all stacks commands
(* process (1, 2, 1) stacks  *)
(*   |> List.iter (fun stack -> *)
(*       List.iter (printf "%c") stack; *)
(*       printf "\n" *)
(*   ) *)
(* ;; *)
