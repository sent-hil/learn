hello = (a, b...)->
  if b
    console.log a, b...
  else
    console.log a

hello "name"
hello "name", {name:"senthil"}
hello "name", {name:"senthil"}, "sdf"
