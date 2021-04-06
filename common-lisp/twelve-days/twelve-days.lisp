(in-package #:cl-user)
(defpackage #:twelve-days
  (:use #:cl)
  (:export #:recite #:print-verse))

(in-package #:twelve-days)

(defparameter *gifts*
  '((12 "Drummers Drumming")
    (11 "Pipers Piping")
    (10 "Lords-a-Leaping")
    (9 "Ladies Dancing")
    (8 "Maids-a-Milking")
    (7 "Swans-a-Swimming")
    (6 "Geese-a-Laying")
    (5 "Gold Rings")
    (4 "Calling Birds")
    (3 "French Hens")
    (2 "Turtle Doves")
    (1 "Partridge in a Pear Tree"))
  "Our list of gifts to iterate.")

(defun recite (&optional begin end)
  "Returns a string of the requested verses for the 12 Days of Christmas."
  (declare (type (or null (integer 1 12)) begin end))
    (format nil "~{~&~A~}"
            (loop for i from (or begin 1) to (or end begin 12)
                  collect (print-verse i))))

(defun print-verse (cur)
  "Returns a string of a specific verse of the 12 Days of Christmas"
  ;; The format line here is a bit write-only, but we could replace with a
  ;; loop or something if we liked.
  (apply #'format
         nil
         "~&On the ~:R day of Christmas my true love gave to me:~
          ~#[ nothing, the ingrate~; a ~{~*~A~}~:;~
          ~@{ ~#[0~;and a ~{~*~A~}~:;~{~R ~A~}~^,~]~}~]."
         cur
         (nthcdr (- 12 cur) *gifts*)))
