(ns day-23.core
  (:require [shams.priority-queue :as pq]))

; #############
; #...........#
; ###B#D#C#A###
;   #D#C#B#A#
;   #D#B#A#C#
;   #C#D#B#A#
;   #########

; #############
; #123456789AB#
; ###.#.#.#.### 1
;   #.#.#.#.#   2
;   #.#.#.#.#   3
;   #.#.#.#.#   4
;   #########

(def edges
  #{
    [:h1 :h2]
    [:h2 :h3]
    [:h3 :h4]
    [:h4 :h5]
    [:h5 :h6]
    [:h6 :h7]
    [:h7 :h8]
    [:h8 :h9]
    [:h9 :hA]
    [:hA :hB]
    [:a4 :a3]
    [:a3 :a2]
    [:a2 :a1]
    [:a1 :h3]
    [:b4 :b3]
    [:b3 :b2]
    [:b2 :b1]
    [:b1 :h5]
    [:c4 :c3]
    [:c3 :c2]
    [:c2 :c1]
    [:c1 :h7]
    [:d4 :d3]
    [:d3 :d2]
    [:d2 :d1]
    [:d1 :h9]
    }
  )

(def nodes
  (set
   (concat
    (for [e edges]
      (get e 0))
    (for [e edges]
      (get e 1)))))

(def neighbors
  (apply
   hash-map
   (apply
    concat
    (for [n nodes]
      [n
       (map (fn [[src dest]] dest)
            (filter
             (fn [[src dest]] (= src n))
             (concat
              (for [e edges]
                e)
              (for [e edges]
                [(get e 1) (get e 0)]))))]))))

(get neighbors :d1)

(def initial-board
  {:energy 0
   :positions {:c_a1 {:position :d1 :state :start}
               :c_a2 {:position :d2 :state :start}
               :c_a3 {:position :c3 :state :start}
               :c_a4 {:position :d4 :state :start}
               :c_b1 {:position :a1 :state :start}
               :c_b2 {:position :c2 :state :start}
               :c_b3 {:position :b3 :state :start}
               :c_b4 {:position :c4 :state :start}
               :c_c1 {:position :c1 :state :start}
               :c_c2 {:position :b2 :state :start}
               :c_c3 {:position :d3 :state :start}
               :c_c4 {:position :a4 :state :start}
               :c_d1 {:position :b1 :state :start}
               :c_d2 {:position :a2 :state :start}
               :c_d3 {:position :a3 :state :start}
               :c_d4 {:position :b4 :state :start}}})

(defn is-node-open? [board node]
  ((complement some)
   true?
   (map (fn [c] (= node (get c :position)))
        (vals (get board :positions)))))

(is-node-open? initial-board :h1)

(defn accessible-nodes
  "returns all nodes we can visit from the origin, along with how many moves it would take to get there"
  [board origin]
  (loop [stack (vector [origin 0])
         visited {} ; includes moves
         ]
    (if (empty? stack)
      visited
      (letfn [(visited? [v] (some #(= % v) (keys visited)))]
        (let [[v moves] (peek stack)
              neighbors (get neighbors v)
              not-occupied (filter #(is-node-open? board %) neighbors)
              not-visited (filter (complement visited?) not-occupied)
              new-stack (into (pop stack) (map (fn [x] [x (inc moves)]) not-visited))]
          (if (visited? v)
            (recur new-stack visited)
            (recur new-stack (assoc visited v moves))))))))

(accessible-nodes initial-board :a2)

(defn energy [creature]
  (case creature
    (:c_a1 :c_a2 :c_a3 :c_a4) 1
    (:c_b1 :c_b2 :c_b3 :c_b4) 10
    (:c_c1 :c_c2 :c_c3 :c_c4) 100
    (:c_d1 :c_d2 :c_d3 :c_d4) 1000))


(defn make-move "Create a new board with the creature moved to the destination, if it's possible"
  [board move]
  (let [creature (get move :creature)
        dest (get move :dest)
        new-state (get move :new-state)
        dist (get move :dist)]
    (-> board
        (assoc-in [:positions creature :position] dest)
        (assoc-in [:positions creature :state] new-state)
        (assoc-in [:energy] (+ (:energy board) (* (energy creature) dist))))))

(def valid-start-destinations
  #{:h1
    :h2
    ; h3 isn't valid, since it's an intersection
    :h4
    ; h5 isn't valid, since it's an intersection
    :h6
    ; h7 isn't valid, since it's an intersection
    :h8
    ; h9 isn't valid, since it's an intersection
    :hA
    :hB
    } ; hallways
  )

(defn get-creatures [board] (keys (get board :positions)))

(defn whats-at [board loc]
  (let [creatures (get-creatures board)
        creature-at-loc (first (filter #(= loc (get-in board [:positions % :position])) creatures))
        ]
    creature-at-loc
  ))

(defn get-creature-state [board creature]
  (if (nil? creature) nil
      (get-in board [:positions creature :state])))

(get-creature-state initial-board :c_a1)

(defn valid-midway-destinations [board creature]
  "There's only ever one valid destination from midway: The home room furthest in."
  (letfn [(undone? [node] (not (= :end (get-creature-state board (whats-at board node)))))]
    #{(case creature
        (:c_a1, :c_a2, :c_a3, :c_a4) (if (undone? :a4) :a4 (if (undone? :a3) :a3 (if (undone? :a2) :a2 :a1)))
        (:c_b1, :c_b2, :c_b3, :c_b4) (if (undone? :b4) :b4 (if (undone? :b3) :b3 (if (undone? :b2) :b2 :b1)))
        (:c_c1, :c_c2, :c_c3, :c_c4) (if (undone? :c4) :c4 (if (undone? :c3) :c3 (if (undone? :c2) :c2 :c1)))
        (:c_d1, :c_d2, :c_d3, :c_d4) (if (undone? :d4) :d4 (if (undone? :d3) :d3 (if (undone? :d2) :d2 :d1))))}))

(valid-midway-destinations initial-board :c_a1)

(defn legal-moves-for-creature [board creature]
  (let [state    (get-in board [:positions creature :state])
        position (get-in board [:positions creature :position])
        valid-nodes (case state
                      :start valid-start-destinations
                      :midway (valid-midway-destinations board creature)
                      :end {}
                      :else {})
        reachable-nodes (accessible-nodes board position)
        legal-destinations (filter (partial contains? valid-nodes) (keys reachable-nodes))
        legal-moves (map (fn [dest]
                           {:creature creature
                            :dest dest
                            :new-state (case state
                                         :start :midway
                                         :midway :end
                                         :end :end)
                            :dist (get reachable-nodes dest)
                            :energy (* (get reachable-nodes dest) (energy creature))})
                         legal-destinations)]
    legal-moves))

(legal-moves-for-creature initial-board :c_d1)

(def board-one-moved
  {:energy 2
   :positions {:c_a1 {:position :h8 :state :midway}
               :c_a2 {:position :d2 :state :start}
               :c_b1 {:position :a1 :state :start}
               :c_b2 {:position :c1 :state :start}
               :c_c1 {:position :c1 :state :start}
               :c_c2 {:position :a2 :state :start}
               :c_d1 {:position :b1 :state :start}
               :c_d2 {:position :b2 :state :start}}})

(valid-midway-destinations board-one-moved :c_a1)

(legal-moves-for-creature board-one-moved :c_a1)

(defn get-legal-moves "Generate all legal moves from a given board"
  [board]
  (let [creatures (keys (get board :positions))
        moves (mapcat (fn [c] (legal-moves-for-creature board c)) creatures)
        boards (map (partial make-move board) moves)
        ]
    boards))

(-> (get-legal-moves initial-board)
    (first))


(defn cmp-boards [x y]
  (let [get-energy #(get % :energy)
        c (compare (get-energy x) (get-energy y))]
    ; (println "comparing" x y ": " c)
    (if (not (= c 0)) c
        1 ; tie-breaker
        )))

(-> (sorted-set-by cmp-boards initial-board board-one-moved)
    (rest)
    (into initial-board)
    )

(-> (sorted-set-by cmp-boards initial-board )
    (rest)
    (into initial-board)
    )

(-> (sorted-set 21  3)
    (rest))

(defn is-winning-board [board]
  (every? (fn [creature] (= :end (get creature :state))) (vals (get board :positions)))
  )

(is-winning-board initial-board)

(def winning-board
  {:energy 9000
   :positions {:c_a1 {:position :h6 :state :end}
               :c_a2 {:position :d2 :state :end}
               :c_b1 {:position :a1 :state :end}
               :c_b2 {:position :c2 :state :end}
               :c_c1 {:position :c1 :state :end}
               :c_c2 {:position :a2 :state :end}
               :c_d1 {:position :b1 :state :end}
               :c_d2 {:position :b2 :state :end}}})

(is-winning-board winning-board)

(defn print-creature [creature]
  (case creature
    nil "·"
    (:c_a1 :c_a2 :c_a3 :c_a4) "A"
    (:c_b1 :c_b2 :c_b3 :c_b4) "B"
    (:c_c1 :c_c2 :c_c3 :c_c4) "C"
    (:c_d1 :c_d2 :c_d3 :c_d4) "D"
    ))


(whats-at initial-board :a1)

(defn print-board [board]
  (apply
   (partial printf " \n%s%s%s%s%s%s%s%s%s%s%s\n  %s %s %s %s\n  %s %s %s %s\n  %s %s %s %s\n  %s %s %s %s\n")
   (map #(print-creature (whats-at board %)) [:h1 :h2 :h3 :h4 :h5 :h6 :h7 :h8 :h9 :hA :hB
                                              :a1 :b1 :c1 :d1
                                              :a2 :b2 :c2 :d2
                                              :a3 :b3 :c3 :d3
                                              :a4 :b4 :c4 :d4
                                              ])))

(print-board initial-board)

(defn visited-value [board]
  (map #(get % :position) (vals (get board :positions))))

(def middle-board
  {:energy 9000
   :positions {:c_a1 {:position :h6 :state :midway}
               :c_a2 {:position :d2 :state :start}
               :c_b1 {:position :a1 :state :start}
               :c_b2 {:position :c2 :state :start}
               :c_c1 {:position :c1 :state :start}
               :c_c2 {:position :a2 :state :start}
               :c_d1 {:position :b1 :state :start}
               :c_d2 {:position :b2 :state :start}}})

(visited-value middle-board)

(defn board-search [initial-board]
  (loop [queue (pq/priority-queue #(- (get % :energy)) :elements [initial-board])
         visited #{}
         its 0]
    (let [board (peek queue)
          visited-val (visited-value board)
          winner? (is-winning-board board)
          ]
      (if (> its 1000000) (throw :foo))
      (if (or winner? (= 0 (mod its 1000)))
        (do
          (println "its " its)
          (println "energy" (get board :energy))
          (println "queue " (count queue))
          ; (println "queue " (map #(get % :energy) queue))
          ; (println board)
          (print-board board)))
      (if winner? (get board :energy)
          (if (contains? visited visited-val) (recur (pop queue) visited (inc its))
          (recur (into (pop queue) (get-legal-moves board)) (conj visited (visited-value board)) (inc its)))))))


(is-node-open? middle-board :h6)
(->> middle-board
     (get-legal-moves)
     (map print-board))

; (board-search initial-board)
