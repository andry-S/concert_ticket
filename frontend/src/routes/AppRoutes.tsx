import { Suspense, lazy, useEffect, useState } from 'react';
import { Routes, Route, Navigate, useLocation } from 'react-router-dom';
import { AnimatePresence } from 'framer-motion';

// Lazy load halaman
const Home = lazy(() => import('../pages/Home'));
const Login = lazy(() => import('../pages/Login'));
const Dashboard = lazy(() => import('../pages/Dashboard'));
const Book = lazy(() => import('../pages/Book'));

const AppRoutes = () => {
  const [userId, setUserId] = useState<string | null>(null);
  const location = useLocation();

  useEffect(() => {
    const storedUserId = localStorage.getItem('user_id');
    setUserId(storedUserId);
  }, [location.pathname]);

  return (
    <Suspense fallback={<div>Loading...</div>}>
      <AnimatePresence mode="wait" initial={false}>
        {/* Gunakan location dan key agar transisi bisa bekerja saat berpindah route */}
        <Routes location={location} key={location.pathname}>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route
            path="/dashboard"
            element={userId ? <Dashboard /> : <Navigate to="/login" />}
          />
          <Route
            path="/book/:id"
            element={userId ? <Book /> : <Navigate to="/login" />}
          />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </AnimatePresence>
    </Suspense>
  );
};

export default AppRoutes;
