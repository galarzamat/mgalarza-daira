<h1 class="text-center text-3xl bg-gray-300 font-bold pb-1 mb-2">Calculadora Básica</h1>
<div class="bg-black py-7 text-white m-2">
    <p></p>
</div>
<div class="m-2">
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-orange flex-1">CE</button>
        <button class="btn btn-orange flex-1">C</button>
        <button class="btn btn-red flex-1">±</button>
        <button class="btn btn-red flex-1">%</button>
    </div>
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue flex-1">7</button>
        <button class="btn btn-blue flex-1">8</button>
        <button class="btn btn-blue flex-1">9</button>
        <button class="btn btn-red flex-1">÷</button>
    </div>
    
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue flex-1">4</button>
        <button class="btn btn-blue flex-1">5</button>
        <button class="btn btn-blue flex-1">6</button>
        <button class="btn btn-red flex-1">x</button>
    </div>
    
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue flex-1">1</button>
        <button class="btn btn-blue flex-1">2</button>
        <button class="btn btn-blue flex-1">3</button>
        <button class="btn btn-red flex-1">-</button>
    </div>
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue flex-1">0</button>
        <button class="btn btn-blue flex-1">.</button>
        <button class="btn btn-green flex-1">=</button>
        <button class="btn btn-red flex-1">+</button>
    </div>
</div>
<style>
   .btn{
    @apply flex-grow text-center p-4; /* Asegura que cada botón se expanda para llenar el espacio asignado */
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
