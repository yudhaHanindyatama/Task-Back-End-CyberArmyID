-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jun 17, 2024 at 05:33 AM
-- Server version: 8.0.30
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_mahasiswa`
--

-- --------------------------------------------------------

--
-- Table structure for table `mahasiswa`
--

CREATE TABLE `mahasiswa` (
  `id_mahasiswa` int NOT NULL,
  `nama` varchar(50) NOT NULL,
  `kelas` varchar(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `mahasiswa`
--

INSERT INTO `mahasiswa` (`id_mahasiswa`, `nama`, `kelas`) VALUES
(1, 'Satria', '3A'),
(2, 'Rio', '3A'),
(3, 'Santi', '3C'),
(4, 'Rani', '3B'),
(5, 'Andi', '3A'),
(6, 'Yuni', '3B'),
(7, 'Vera', '3C');

-- --------------------------------------------------------

--
-- Table structure for table `nilai_mahasiswa`
--

CREATE TABLE `nilai_mahasiswa` (
  `id_nilai` int NOT NULL,
  `id_mahasiswa` int NOT NULL,
  `matematika` int NOT NULL,
  `kimia` int NOT NULL,
  `fisika` int NOT NULL,
  `biologi` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `nilai_mahasiswa`
--

INSERT INTO `nilai_mahasiswa` (`id_nilai`, `id_mahasiswa`, `matematika`, `kimia`, `fisika`, `biologi`) VALUES
(1, 1, 90, 88, 100, 82),
(2, 2, 87, 86, 85, 84),
(3, 3, 60, 86, 54, 75),
(4, 4, 80, 65, 76, 78),
(5, 5, 65, 62, 75, 65),
(6, 6, 74, 73, 72, 71),
(7, 7, 70, 65, 64, 63);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `mahasiswa`
--
ALTER TABLE `mahasiswa`
  ADD PRIMARY KEY (`id_mahasiswa`);

--
-- Indexes for table `nilai_mahasiswa`
--
ALTER TABLE `nilai_mahasiswa`
  ADD PRIMARY KEY (`id_nilai`),
  ADD KEY `id_mahasiswa` (`id_mahasiswa`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `mahasiswa`
--
ALTER TABLE `mahasiswa`
  MODIFY `id_mahasiswa` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `nilai_mahasiswa`
--
ALTER TABLE `nilai_mahasiswa`
  MODIFY `id_nilai` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `nilai_mahasiswa`
--
ALTER TABLE `nilai_mahasiswa`
  ADD CONSTRAINT `nilai_mahasiswa_ibfk_1` FOREIGN KEY (`id_mahasiswa`) REFERENCES `mahasiswa` (`id_mahasiswa`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
