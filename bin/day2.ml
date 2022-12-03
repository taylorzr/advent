open Format
open Advent

let n = Str.regexp "\n";;
let s = Str.regexp " ";;

let item_values = Hashtbl.create 3;;
Hashtbl.add item_values "rock" 1;;
Hashtbl.add item_values "paper" 2;;
Hashtbl.add item_values "scissors" 3;;

let lose_to = Hashtbl.create 3;;
Hashtbl.add lose_to "rock" "scissors";;
Hashtbl.add lose_to "paper" "rock";;
Hashtbl.add lose_to "scissors" "paper";;

let win_against = Hashtbl.create 3;;
Hashtbl.add win_against "rock" "paper";;
Hashtbl.add win_against "paper" "scissors";;
Hashtbl.add win_against "scissors" "rock";;

let load input =
  Str.split n input
  |> List.map (fun game ->
      Str.split s game
      |> List.map(fun item ->
          match item with
          | "A" | "X" -> "rock"
          | "B" | "Y" -> "paper"
          | "C" | "Z" -> "scissors"
          | _ -> failwith "oops"
      )
  )
;;

let score game =
  let x = match game with
  | [_; b] -> Hashtbl.find item_values b
  | _ -> 0 in
  let y = match game with
  | [a; b] when a == b -> 3
  | ["scissors"; "rock"] -> 6
  | ["rock"; "paper"] -> 6
  | ["paper"; "scissors"] -> 6
  | _ -> 0 in
  (* printf "%d + %d\n" x y;  *)
  x + y
;;

let load2 input =
  Str.split n input
  |> List.map (fun game ->
      Str.split s game
      |> List.map(fun item ->
          match item with
          | "A" ->  "rock"
          | "B" ->  "paper"
          | "C" ->  "scissors"
          | "X" ->  "lose"
          | "Y" ->  "draw"
          | "Z" ->  "win"
          | _ -> failwith "oops"
      )
  )
;;

let score2 game =
  match game with
  | [a; "draw"] -> 3 + Hashtbl.find item_values a
  | [a; "lose"] -> 0 + Hashtbl.find item_values (Hashtbl.find lose_to a)
  | [a; "win"] -> 6 + Hashtbl.find item_values (Hashtbl.find win_against a)
  | _ -> 0
;;

load d2sample |> List.map score |> List.fold_left (+) 0 |> printf "%d\n";;
load d2input |> List.map score |> List.fold_left (+) 0 |> printf "%d\n";;
load2 d2sample |> List.map score2 |> List.fold_left (+) 0 |> printf "%d\n";;
load2 d2input |> List.map score2 |> List.fold_left (+) 0 |> printf "%d\n";;
