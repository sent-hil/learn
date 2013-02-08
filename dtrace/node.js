new Error().stack;

function main() { func1(); }
function func1() { func2(); }

function func2()
{
  (function () {
      for (;;)
        ;
    })();
}

main();
