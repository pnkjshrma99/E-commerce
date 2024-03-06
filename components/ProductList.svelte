<!-- components/ProductList.svelte -->
<script>
    import { onMount } from 'svelte';
  
    export let isLoggedIn;
    export let addToCart;
  
    let products = [];
  
    onMount(() => {
      fetchProducts();
    });
  
    async function fetchProducts() {
      try {
        const response = await fetch('/api/products');
        products = await response.json();
      } catch (error) {
        console.error('Error fetching products:', error);
      }
    }
  </script>
  
  <ul>
    {#each products as product}
      <li>
        <h3>{product.name}</h3>
        <p>{product.description}</p>
        <p>Price: ${product.price}</p>
        <p>Stock: {product.stock}</p>
        {#if isLoggedIn}
          <button on:click={() => addToCart(product)}>Add to Cart</button>
        {/if}
      </li>
    {/each}
  </ul>