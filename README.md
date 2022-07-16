# nfttrader
A platform for NFT traders

### container/dotweb
  web容器：服务绑定dotweb容器引擎
  build an container by dotweb.
- internal 
  提供基础基础方法
  provide basic functions for web container
  读取配置文件
  read configuration
  读取文件/本地资源
  read files
  controller路由解析
  controller route wrapper
- infrastructure
  容器基础组件
  provider basic component for web container.
  socket绑定
  socket binding
  请求队列
  queue
  协程配置
  goroutine configuration
  内容地址绑定
  content address binding
### model
- po
  数据实体
  entities map to table
- bo
  业务实体
  business objects which were built by business model. 
- vo
  视图实体
  provide item to page view.
# dao
  数据获取
  Lay is focus on  data access.
# business
  业务逻辑
  Lay is focus on business logics and interact with BlockChain.
# controller
  controller view
  Bridge a relationship for page views and business lay.
# utility
  通用工具
  common tools.
### images
  图片存储地址
  store images
### content
  站点内容存储地址
  store html js and css
### resources
  配置文件信息存储地址
  application configurations.