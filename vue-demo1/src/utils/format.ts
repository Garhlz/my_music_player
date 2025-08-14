// src/utils/format.ts

import { format, formatDistanceToNow, parseISO } from 'date-fns';
import { zhCN } from 'date-fns/locale';

/**
 * 格式化时长，将秒数转换为 `mm:ss` 格式的字符串。
 * @param seconds - 总秒数
 * @returns 格式化后的字符串，例如 "04:30"
 */
export const formatDuration = (seconds: number | null | undefined): string => {
  // 如果输入无效或为0，返回 '00:00'
  if (!seconds || seconds <= 0) {
    return '00:00';
  }

  const minutes = Math.floor(seconds / 60);
  const remainingSeconds = Math.floor(seconds % 60);

  // 使用 padStart 补零，比自定义函数更现代、更简洁
  const formattedMinutes = String(minutes).padStart(2, '0');
  const formattedSeconds = String(remainingSeconds).padStart(2, '0');

  return `${formattedMinutes}:${formattedSeconds}`;
};

/**
 * 格式化日期，将 Date 对象或时间字符串转换为 `yyyy-MM-dd` 格式。
 * 利用 date-fns 库，代码更简洁且功能更强大。
 * @param date - Date 对象、ISO 时间字符串或时间戳
 * @returns 格式化后的字符串，例如 "2025-08-14"
 */
export const formatDate = (date: Date | string | number | null | undefined): string => {
  if (!date) {
    return '';
  }

  try {
    return format(new Date(date), 'yyyy-MM-dd');
  } catch (error) {
    console.error('日期格式化失败:', error);
    return '无效日期';
  }
};

/**
 * 将 ISO 格式的时间字符串转换为人性化的相对时间。
 * @param timeStr - 从后端获取的时间字符串 (例如 "2025-08-14T12:43:00.123Z")
 * @returns 格式化后的字符串 (例如 "5分钟前")
 */
export const formatRelativeTime = (timeStr?: string): string => {
  if (!timeStr) {
    return '';
  }

  try {
    // parseISO 会将 ISO 格式的字符串转换为 Date 对象
    const date = parseISO(timeStr);

    // formatDistanceToNow 计算该日期到现在的相对时间
    return formatDistanceToNow(date, {
      addSuffix: true, // 添加 "前" 或 "后" 的后缀
      locale: zhCN,      // 使用中文语言包
    });
  } catch (error) {
    console.error('相对时间格式化失败:', error);
    return timeStr; // 如果格式化失败，返回原始字符串
  }
};

// 注意：你原来的 formatNumber 函数现在已经不再需要，因为它已经被更现代的方法所取代。