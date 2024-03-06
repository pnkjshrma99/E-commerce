<!-- App.svelte -->
<script>
    import ProductList from './components/ProductList.svelte';
    import Cart from './components/Cart.svelte';
    import LoginForm from './components/LoginForm.svelte';
    import RegisterForm from './components/RegisterForm.svelte';
  
    let isLoggedIn = false;
    let cart = [];
  
    function handleLogin() {
      isLoggedIn = true;
    }
  
    function handleLogout() {
      isLoggedIn = false;
      cart = [];
    }
  
    function addToCart(product) {
      cart = [...cart, product];
    }
  
    function removeFromCart(index) {
      cart = cart.filter((_, i) => i !== index);
    }
  </script>
  
  <main>
    <header>
      <h1>E-Commerce Platform</h1>
      <nav>
        {#if isLoggedIn}
          <button on:click={handleLogout}>Logout</button>
          <Cart {cart} {removeFromCart} />
        {:else}
          <LoginForm on:login={handleLogin} />
          <RegisterForm />
        {/if}
      </nav>
    </header>
  
    <section>
      <ProductList {isLoggedIn} {addToCart} />
    </section>
  </main>