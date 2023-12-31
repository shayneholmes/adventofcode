(defproject day-24 "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :license {:name "EPL-2.0 OR GPL-2.0-or-later WITH Classpath-exception-2.0"
            :url "https://www.eclipse.org/legal/epl-2.0/"}
  :dependencies [[org.clojure/clojure "1.10.3"]]
  :repl-options {:init-ns day-24.core})

(defmacro get-val [state v]
  `(if (number? ~v) ~v (get ~state ~v)))
(defmacro inp [state v]
  `(assoc ~state
          ~v (first (get ~state :input))
          :input (rest (get ~state :input))))
(defmacro add [state a b]
  `(assoc ~state ~a (+ (get-val ~state ~a) (get-val ~state ~b))))
(defmacro mul [state a b]
  `(assoc ~state ~a (* (get-val ~state ~a) (get-val ~state ~b))))
(defmacro mod [state a b]
  `(assoc ~state ~a (clojure.core/mod (get-val ~state ~a) (get-val ~state ~b))))
(defmacro div [state a b]
  `(assoc ~state ~a (unchecked-divide-int (get-val ~state ~a) (get-val ~state ~b))))
(defmacro eql [state a b]
  `(assoc ~state ~a (if (clojure.core/= (get-val ~state ~a) (get-val ~state ~b)) 1 0)))

(def input '(1 2 3 4 5 6 7 8 9 1 2 3 4 5))

(let [w :w
      x :x
      y :y
      z :z
      state {w 0
             x 0
             y 0
             z 0
             :input input}]
  (-> state
      (inp w)
      (inp w)
      (mul w w)
      (inp x)
      (eql x x)
      (eql x x)
      ; (mul x 0)
      ; (add x z)
      ; (mod x 26)
      ; (div z 1)
      ; (add x 12)
      (identity)))

(let [w :w
      x :x
      y :y
      z :z
      state {w 0
             x 0
             y 0
             z 0
             :input input}]
  (-> state
      (inp w) ; w = in[0]
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 1)
      (add x 12)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 6)
      (mul y x)
      (add z y)
      (inp w) ; w = in[1]
      (mul x 0) ; x =0
      (add x z)
      (mod x 26)
      (div z 1)
      (add x 10)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 6)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 1)
      (add x 13)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 3)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 26)
      (add x -11)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 11)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 1)
      (add x 13)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 9)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 26)
      (add x -1)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 3)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 1)
      (add x 10)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 13)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 1)
      (add x 11)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 6)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 26)
      (add x 0)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 14)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 1)
      (add x 10)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 10)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 26)
      (add x -5)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 12)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 26)
      (add x -16)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0)
      (add y w)
      (add y 10)
      (mul y x)
      (add z y)
      (inp w)
      (mul x 0)
      (add x z)
      (mod x 26)
      (div z 26)
      (add x -7)
      (eql x w)
      (eql x 0) ; 1 or 0
      (mul y 0)
      (add y 25)
      (mul y x)
      (add y 1)
      (mul z y)
      (mul y 0) ;y = 0
      (add y w)
      (add y 11)
      (mul y x)
      (add z y)
      (inp w) ; w = last digit
      (mul x 0) ; x = 0
      (add x z) ; x = z
      (mod x 26) ; x %= 26
      (div z 26) ; z /= 26, must equal 0, so z < 26
      (add x -11) ; x -= 11
      (eql x w) ; 1 or 0 must equal 1, so x == w (last digit))
      (eql x 0) ; invert (1<>0) , must equal 0
      (mul y 0) ; y =0
      (add y 25) ;
      (mul y x) ;
      (add y 1) ;
      (mul z y) ; z = 0 | y = 0, and y can't equal 0 because of +1 above, so z = 0
      (mul y 0) ; y = 0; any references to y before this can be optimized away?
      (add y w)
      (add y 15)
      (mul y x) ; x = 0 | y = 0, and y can't equal 0 because of +15 above, so x = 0
      (add z y) ; z = 0 & y = 0

      (identity)))


; '(inp t t)
; (macroexpand-1 '(inp t t))
; (macroexpand '(add w 25 (w)))

; (rest input)
; (inp x (add x 5 x))
; (dir user)

; (let [t :t]
;   (-> {:input '(1 2 3 4 5)}
;       (inp t)
;       (add t 5)
;       (identity )))

; (macroexpand '(inp x x))
; (macroexpand '(start input (inp x x)))
; (start input (inp x x))
; (macroexpand '(let [input '(1 2 3 4 5)
;       a (first input)
;       input (rest input)
;       b ( first input)] [a b]))
