<!-- components/RegisterForm.svelte -->
<script>
    let name = '';
    let email = '';
    let password = '';
  
    async function handleRegister() {
      try {
        const response = await fetch('/api/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ name, email, password }),
        });
  
        if (response.ok) {
          console.log('Registration successful');
          // Handle success, e.g., redirect to login page
        } else {
          console.error('Registration failed');
        }
      } catch (error) {
        console.error('Error during registration:', error);
      }
    }
  </script>
  
  <div>
    <h2>Register</h2>
    <form on:submit|preventDefault={handleRegister}>
      <input type="text" placeholder="Name" bind:value={name} required />
      <input type="email" placeholder="Email" bind:value={email} required />
      <input type="password" placeholder="Password" bind:value={password} required />
      <button type="submit">Register</button>
    </form>
  </div>