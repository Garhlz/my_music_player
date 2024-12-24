// 格式化时长（将秒数转换为 mm:ss 格式）
export const formatDuration = (duration) => {
  if (!duration) return '00:00'
  
  const minutes = Math.floor(duration / 60)
  const seconds = Math.floor(duration % 60)
  
  const formatNumber = (num) => {
    return num < 10 ? `0${num}` : num
  }
  
  return `${formatNumber(minutes)}:${formatNumber(seconds)}`
}

// 格式化日期
export const formatDate = (date) => {
  if (!date) return ''
  
  const d = new Date(date)
  const year = d.getFullYear()
  const month = formatNumber(d.getMonth() + 1)
  const day = formatNumber(d.getDate())
  
  return `${year}-${month}-${day}`
}

// 格式化数字（如播放次数）
export const formatNumber = (num) => {
  if (num < 10) return `0${num}`
  return num.toString()
} 