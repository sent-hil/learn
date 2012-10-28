(define fact
  (lambda (n)
    (if (= n 0)
      1        ;Base case: return 1
      (* n (fact (- n 1))))))
