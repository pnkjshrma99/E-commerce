<!-- components/LoginForm.svelte -->
<script>
    import { createEventDispatcher } from 'svelte';
  
    const dispatch = createEventDispatcher();
  
    let email = '';
    let password = '';
  
    async function handleLogin() {
      try {
        const response = await fetch('/api/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ email, password }),
        });
  
        if (response.ok) {
          dispatch('login');
        } else {
          console.error('Login failed');
        }
      } catch (error) {
        console.error('Error during login:', error);
      }
    }
  </script>
  
  <div>
    <h2>Login</h2>
    <form on:submit|preventDefault={handleLogin}>
      <input type="email" placeholder="Email" bind:value={email} required />
      <input type="password" placeholder="Password" bind:value={password} required />
      <button type="submit">Login</button>
    </form>
  </div>