type nucleotide = A | C | G | T

let hamming_distance l1 l2 = match l1, l2 with
    [], [] -> Ok 0
  | _::_, [] -> Error "right strand must not be empty"
  | [], _::_ -> Error "left strand must not be empty"
  | l1, l2 -> let rec helper ll rl acc = match ll, rl with
      | h1::t1, h2::t2 when h1=h2 -> helper t1 t2 acc
      | _::t1, _::t2 -> helper t1 t2 (acc+1)
      | [], _::_ -> Error "left and right strands must be of equal length"
      | _::_, [] -> Error "left and right strands must be of equal length"
      | [], [] -> Ok acc
    in helper l1 l2 0
