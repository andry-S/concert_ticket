-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jun 24, 2025 at 08:47 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `user_service`
--

-- --------------------------------------------------------

--
-- Table structure for table `bookings`
--

CREATE TABLE `bookings` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `event_id` int(11) NOT NULL,
  `total_price` int(11) NOT NULL,
  `ticket_quantity` int(11) NOT NULL,
  `status` varchar(20) DEFAULT 'waiting_payment',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `payment_id` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `bookings`
--

INSERT INTO `bookings` (`id`, `user_id`, `event_id`, `total_price`, `ticket_quantity`, `status`, `created_at`, `updated_at`, `payment_id`) VALUES
(1, 1, 101, 500000, 2, 'pending', '2025-06-23 11:54:00', '2025-06-23 11:54:00', NULL),
(2, 2, 102, 250000, 1, 'paid', '2025-06-23 11:54:00', '2025-06-23 11:54:00', NULL),
(3, 3, 103, 1000000, 4, 'cancelled', '2025-06-23 11:54:00', '2025-06-23 11:54:00', NULL),
(4, 1, 104, 750000, 3, 'paid', '2025-06-23 11:54:00', '2025-06-23 11:54:00', NULL),
(5, 4, 101, 250000, 1, 'pending', '2025-06-23 11:54:00', '2025-06-23 11:54:00', NULL),
(6, 2, 103, 500000, 2, 'paid', '2025-06-23 11:54:00', '2025-06-23 11:54:00', NULL),
(7, 1, 2, 750000, 3, 'pending', '2025-06-23 21:28:44', '2025-06-23 21:28:44', NULL),
(8, 3, 3, 250000, 1, 'pending', '2025-06-23 21:48:49', '2025-06-23 21:48:49', NULL),
(9, 3, 1, 250000, 1, 'pending', '2025-06-23 21:49:08', '2025-06-23 21:49:08', NULL),
(10, 3, 1, 500000, 2, 'pending', '2025-06-23 22:01:05', '2025-06-23 22:01:05', NULL),
(11, 3, 2, 250000, 1, 'pending', '2025-06-23 22:01:58', '2025-06-23 22:01:58', NULL),
(12, 3, 2, 250000, 1, 'pending', '2025-06-23 22:03:04', '2025-06-23 22:03:04', NULL),
(13, 3, 1, 250000, 1, 'pending', '2025-06-24 09:49:41', '2025-06-24 09:49:41', NULL),
(14, 3, 1, 250000, 1, 'pending', '2025-06-24 10:00:49', '2025-06-24 10:00:49', NULL),
(15, 3, 4, 250000, 1, 'pending', '2025-06-24 10:24:30', '2025-06-24 10:24:30', NULL),
(16, 3, 10, 250000, 1, 'pending', '2025-06-24 10:25:18', '2025-06-24 10:25:18', NULL),
(17, 3, 2, 250000, 1, 'pending', '2025-06-24 10:54:14', '2025-06-24 10:54:14', NULL),
(18, 3, 7, 250000, 1, 'pending', '2025-06-24 12:29:13', '2025-06-24 12:29:13', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `payments`
--

CREATE TABLE `payments` (
  `id` int(11) NOT NULL,
  `booking_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `status` varchar(20) DEFAULT 'pending',
  `payment_id` varchar(100) DEFAULT NULL,
  `payment_url` text DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`) VALUES
(3, 'admin', 'admin@mail.com', '$2a$10$UBnEY7iMOwmfXZ8udKzdi.iVARStShek/eYN4K3CsKLOP7N89ftue'),
(4, 'user1', 'user1@mail.com', '$2a$10$UBnEY7iMOwmfXZ8udKzdi.iVARStShek/eYN4K3CsKLOP7N89ftue'),
(5, NULL, 'user2@mail.com', '$2a$10$UBnEY7iMOwmfXZ8udKzdi.iVARStShek/eYN4K3CsKLOP7N89ftue'),
(7, NULL, 'andri@mail.com', '$2a$10$ZszLZtjRE7W2SXaazhiNXeAyUnvNQ/MOvM4EdQf12OMoDWAJ1cCja');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `bookings`
--
ALTER TABLE `bookings`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `payments`
--
ALTER TABLE `payments`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `bookings`
--
ALTER TABLE `bookings`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `payments`
--
ALTER TABLE `payments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
