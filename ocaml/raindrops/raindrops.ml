let raindrop i =
  match i mod 3, i mod 5, i mod 7 with
    0, 0, 0 -> "PlingPlangPlong"
  | 0, 0, _ -> "PlingPlang"
  | 0, _, 0 -> "PlingPlong"
  | _, 0, 0 -> "PlangPlong"
  | 0, _, _ -> "Pling"
  | _, 0, _ -> "Plang"
  | _, _, 0 -> "Plong"
  | _       -> string_of_int i
