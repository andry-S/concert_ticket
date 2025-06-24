import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../styles/main.css'; 

interface LoginForm {
  email: string;
  password: string;
}

const Login = () => {
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<LoginForm>();

  const onSubmit = async (data: LoginForm) => {
    try {
      const res = await axios.post('http://localhost:8001/login', data);


      localStorage.setItem('user_id', res.data.user_id);
      localStorage.setItem('user_name', res.data.user_name);
      localStorage.setItem('user_email', res.data.user_email);

      navigate('/dashboard');
    } catch (err: any) {
      alert('Login gagal');
      if (axios.isAxiosError(err)) {
        console.error('Login error:', err.response?.data);
      } else {
        console.error('Unexpected error:', err);
      }
    }
  };

  return (
    <div className="login-container">
      <form onSubmit={handleSubmit(onSubmit)} className="login-form">
        <h2>Login</h2>

        <input
          {...register('email', { required: 'Email wajib diisi' })}
          type="email"
          placeholder="Email"
        />
        {errors.email && <p className="error-message">{errors.email.message}</p>}

        <input
          {...register('password', { required: 'Password wajib diisi' })}
          type="password"
          placeholder="Password"
        />
        {errors.password && <p className="error-message">{errors.password.message}</p>}

        <button type="submit" disabled={isSubmitting}>
          {isSubmitting ? 'Loading...' : 'Login'}
        </button>
      </form>
    </div>
  );
};

export default Login;
