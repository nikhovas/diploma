import sys


file = sys.argv[1]
to_change_patterns = sys.argv[2:]

data = open(file, 'r').read()
for replace_str in to_change_patterns:
    replace_str_import = replace_str.replace('_', '__')
    find = 'import %s_pb2 as %s__pb2' % (replace_str, replace_str_import)
    replace = 'from proto.%s import %s_pb2 as %s__pb2' % (replace_str, replace_str, replace_str_import)
    data = data.replace(find, replace)

open(file, 'w').write(data)