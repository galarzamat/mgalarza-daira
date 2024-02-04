<script>
    let display='';
    let clear_display= ()=>display='';
    let num1,num2;
    let op;
    function addDisplay(value) {
      display+=value;  
    }
</script>
<div class="bg-black py-7 text-white m-2">
    <p>{display}</p>
</div>
<div class="m-2">
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-orange" on:click={() => clear_display()}>CE</button>
        <button class="btn btn-orange">C</button>
        <button class="btn btn-red">±</button>
        <button class="btn btn-red" >%</button>
    </div>
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue"on:click={() => addDisplay(7)}>7</button>
        <button class="btn btn-blue"on:click={() => addDisplay(8)}>8</button>
        <button class="btn btn-blue"on:click={() => addDisplay(9)}>9</button>
        <button class="btn btn-red" on:click={() => addDisplay("÷")}>÷</button>
    </div>
    
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue"on:click={() => addDisplay(4)}>4</button>
        <button class="btn btn-blue"on:click={() => addDisplay(5)}>5</button>
        <button class="btn btn-blue"on:click={() => addDisplay(6)}>6</button>
        <button class="btn btn-red" on:click={() => addDisplay("x")}>x</button>
    </div>
    
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue"on:click={() => addDisplay(1)}>1</button>
        <button class="btn btn-blue"on:click={() => addDisplay(2)}>2</button>
        <button class="btn btn-blue"on:click={() => addDisplay(3)}>3</button>
        <button class="btn btn-red" on:click={() => addDisplay("-")}>-</button>
    </div>
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue"on:click={() => addDisplay(0)}>0</button>
        <button class="btn btn-blue" on:click={() => addDisplay(".")}>.</button>
        <button class="btn btn-green">=</button>
        <button class="btn btn-red" on:click={() => addDisplay("+")}>+</button>
    </div>
</div>
<style>
   .btn{
    @apply flex-grow text-center p-4 flex-1; /* Asegura que cada botón se expanda para llenar el espacio asignado */
   } 
.btn-blue{
    @apply bg-blue-700 text-white;
}
.btn-orange{
    @apply bg-orange-400 text-white;
}
.btn-green{
    @apply bg-green-600 text-white;
}
.btn-red{
    @apply bg-red-800 text-white;
}
</style>
