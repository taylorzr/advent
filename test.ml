
let rec last l =
  match l with
  | [] -> None
  | [x] -> Some x
  | h :: t -> last t
;;
last ["a" ; "b" ; "c" ; "d"];;

let rec last2 l =
  match l with
  | [] -> None
  | [x; y] -> Some (x, y)
  | h :: t -> last t
;;
last2 ["a" ; "b" ; "c" ; "d"];;

let rec at l i =
  match l with
  | [] -> None
  | h :: t -> if i = 0 then Some h else at t (i - 1)
;;
at ["a" ; "b" ; "c" ; "d"] 2;;

let rec length l =
  match l with
  | [] -> 0
  | _ :: t -> 1 + length t
;;
length ["a"; "b"; "c"];;

