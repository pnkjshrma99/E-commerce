<!-- components/Cart.svelte -->
<script>
    export let cart;
    export let removeFromCart;
  
    let total = 0;
  
    $: {
      total = cart.reduce((acc, product) => acc + product.price, 0);
    }
  </script>
  
  <div>
    <h2>Cart</h2>
    <ul>
      {#each cart as product, index}
        <li>
          {product.name} - ${product.price}
          <button on:click={() => removeFromCart(index)}>Remove</button>
        </li>
      {/each}
    </ul>
    <p>Total: ${total}</p>
    <button>Checkout</button>
  </div>