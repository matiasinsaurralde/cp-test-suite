function MyPreMiddleware(request, session, spec)
  tyk.req.set_header("myluaheader", "myluavalue")
  return request, session
end
