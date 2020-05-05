type dna = [ `A | `C | `G | `T ]
type rna = [ `A | `C | `G | `U ]

let amino_to_amino amino =
  match amino with
  | `A -> `U
  | `C -> `G
  | `G -> `C
  | `T -> `A

let rec to_rna aminos =
  match aminos with
  | [] -> []
  | head::tail -> (amino_to_amino head)::(to_rna tail)
