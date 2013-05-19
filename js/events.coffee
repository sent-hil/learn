class Hello extends EventsEmitter
  run: ->
    "hello"

h = new Hello
console.log h._events
