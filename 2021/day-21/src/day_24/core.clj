(ns day-24.core)

(defn foo
  "I don't do a whole lot."
  [x]
  (println x "Hello, World!"))

(defn get-die []
  (let [counter (atom 100)]
    (fn [] (swap! counter (fn [old] (-> old (mod 100) (inc)))))))

(defn play-game []
  (let [die (get-die)]
    (loop [p1 [7 0] ; pos, score
           p2 [9 0]
           p1-turn true
           roll-count 0
           ]
      ; take a turn
      (let [roll (+ (die) (die) (die))
            [pos score] (if p1-turn p1 p2)
            [other-pos other-score] (if p1-turn p2 p1)
            new-pos-premod (mod (+ pos roll) 10)
            new-pos (if (= new-pos-premod 0) 10 new-pos-premod)
            new-score (+ score new-pos)
            new-p1 (if p1-turn [new-pos new-score] p1)
            new-p2 (if p1-turn p2 [new-pos new-score])
            new-turn (not p1-turn)
            new-roll-count (+ 3 roll-count)
            ]
        (println "player" (if p1-turn "1" "2") "rolled" roll "landed on" new-pos "score is" new-score)
        (if (>= new-score 1000)
          (* other-score new-roll-count) ; done!
          (recur new-p1 new-p2 new-turn new-roll-count)
          )
        ))))

; (play-game)

(defn play-quantum-game []
  (let [rolls [
               {:value 3 :incidence 1}
               {:value 4 :incidence 3}
               {:value 5 :incidence 6}
               {:value 6 :incidence 7}
               {:value 7 :incidence 6}
               {:value 8 :incidence 3}
               {:value 9 :incidence 1}
               ]]

    (loop [states (for [roll rolls]
                    {:p1 [7 0]
                     :p2 [9 0]
                     :p1-turn true
                     :rolled (get roll :value)
                     :incidence (get roll :incidence)})
           p1score 0
           p2score 0
           iterations 0
           ]
      ; (println (first states))
      (if (= 0 (mod iterations 1000000)) (println iterations "iterations, " p1score " vs " p2score ", " (count states) "states queued"))
      (if (empty? states) [p1score p2score]
          ; take a turn

          (let [state (first states)
                p1 (get state :p1)
                p2 (get state :p2)
                p1-turn (get state :p1-turn)
                roll (get state :rolled)
                incidence (get state :incidence)
                [pos score] (if p1-turn p1 p2)
                [other-pos other-score] (if p1-turn p2 p1)
                new-pos-premod (mod (+ pos roll) 10)
                new-pos (if (= new-pos-premod 0) 10 new-pos-premod)
                new-score (+ score new-pos)
                new-p1 (if p1-turn [new-pos new-score] p1)
                new-p2 (if p1-turn p2 [new-pos new-score])
                new-turn (not p1-turn)]
            ; (println "queue has" (count states) "player" (if p1-turn "1" "2") "rolled" roll "landed on" new-pos "score is" new-score)
            (if (>= new-score 21)
              (recur (rest states) (if p1-turn (+ incidence p1score) p1score) (if p1-turn p2score (+ incidence p2score)) (inc iterations)) ; won! tally it
              (let [new-states (for [roll rolls]
                                 {:p1 new-p1
                                  :p2 new-p2
                                  :p1-turn new-turn
                                  :rolled (get roll :value)
                                  :incidence (* incidence (get roll :incidence))})]
                (recur (apply (partial conj (rest states)) new-states) p1score p2score (inc iterations)))))))))

(for [roll1 (range 1 4)
      roll2 (range 1 4)
      roll3 (range 1 4)]
  (+ roll1 roll2 roll3))

; (play-quantum-game)
