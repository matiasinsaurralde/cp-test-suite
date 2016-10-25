from tyk.decorators import *
from gateway import TykGateway as tyk

@Hook
def MyPreMiddleware(request, session, spec):
    request.add_header('mypyheader', 'mypyvalue')
    return request, session
