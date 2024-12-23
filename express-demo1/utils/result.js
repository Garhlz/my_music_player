class Result {
    constructor(code, msg, data = null) {
      this.code = code; // 状态码：1成功，其他为失败
      this.msg = msg;   // 错误或成功信息
      this.data = data; // 返回的数据
    }
  
    // 成功的响应，data 默认为空
    static success() {
      return new Result(1, 'Success');
    }
  
    // 成功的响应，带有数据
    static successWithData(data) {
      return new Result(1, 'Success', data);
    }
  
    // 错误的响应，带有错误消息
    static error(msg) {
      return new Result(0, msg);
    }
}
  
module.exports = Result;
  