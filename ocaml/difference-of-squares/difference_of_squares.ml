(* This is the triangle number formula, squared *)
let square_of_sum n =
  let sum = (n*n+n)/2 in
    sum*sum

(* square pyramidal number formula *)
let sum_of_squares n =
  (n*(n+1)*(2*n+1))/6

let difference_of_squares n =
    square_of_sum n - sum_of_squares n
