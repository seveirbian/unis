#!/usr/bin/env python
# -*- coding: utf-8 -*-
#用于规定字符编码，想要中文正常最好就加上这句

import os.path
import tornado.auth
import tornado.escape
import tornado.httpserver
import tornado.ioloop
import tornado.options
import tornado.web
from tornado.options import define, options
import tornado.httpclient
from tornado.web import HTTPError, asynchronous
from tornado.httpclient import HTTPRequest
from tornado.httpclient import HTTPClient
try:
    from tornado.curl_httpclient import CurlAsyncHTTPClient as AsyncHTTPClient
except ImportError:
    from tornado.simple_httpclient import SimpleAsyncHTTPClient as AsyncHTTPClient
import requests
import logging

define("port", default=8000, help="run on the given port", type=int)
cur_path = os.path.dirname(__file__)
master_url = 'http://202.114.10.146:9898'

class Application(tornado.web.Application):
	def __init__(self):
		handlers = [
		(r"/", MainHandler),
		(r"/(.*\.html)", tornado.web.StaticFileHandler, {"path": os.path.join(cur_path, "Theme")}),
		(r"/.+", ProxyHandler),
		]

		settings = dict(
		template_path=os.path.join(cur_path, "Theme"),
		static_path=os.path.join(cur_path, "Theme/assets"),
		static_url_prefix='/assets/',
		debug=True,
		)

		tornado.web.Application.__init__(self, handlers, **settings)

class MainHandler(tornado.web.RequestHandler):
	def get(self):
		self.render("index.html",)

class ProxyHandler(tornado.web.RequestHandler):
	def get(self):
		url = master_url + self.request.uri
		headers = dict(self.request.headers)
		headers['Host'] = '202.114.10.146'

		try:
			response = HTTPClient().fetch(HTTPRequest(
					url=url,
					method="GET",
					headers=headers,
					follow_redirects=False))
			self._on_proxy(response)
		except tornado.httpclient.HTTPError, x:
			if hasattr(x, "response") and x.response:
				self._on_proxy(x.response)
			else:
				logging.error("Tornado signalled HTTPError %s", x)

	def post(self):
		url = master_url + self.request.uri
		headers = dict(self.request.headers)
		headers['Host'] = '202.114.10.146'

		try:
			response = HTTPClient().fetch(
				HTTPRequest(
					url=url,
					method="POST",
					body=self.request.body,
					headers=headers,
					request_timeout=8,
					follow_redirects=False))
			self._on_proxy(response)
		except tornado.httpclient.HTTPError, x:
			if hasattr(x, "response") and x.response:
				self._on_proxy(x.response)
			else:
				logging.error("Tornado signalled HTTPError %s", x)
				raise HTTPError(501)

	def _on_proxy(self, response):
		if response.error and not isinstance(response.error, 
			tornado.httpclient.HTTPError):
			raise HTTPError(500)
		else:
			self.set_status(response.code)
			for header in ("Date", "Cache-Control", "Server", "Content-Type", "Location"):
				v = response.headers.get(header)
				print header, v
				if v:
					self.set_header(header, v)
			if response.body:
				print response.body
				self.write(response.body)
			self.finish()

def main():
	tornado.options.parse_command_line()
	http_server = tornado.httpserver.HTTPServer(Application())
	http_server.listen(options.port)
	tornado.ioloop.IOLoop.instance().start()

if __name__ == "__main__":
	main()
