local AdminName(name='<untitled>') = 'user_admin_' + name;

{
  host: 'localhost:8080',
  admins: [AdminName(name='Bob'), AdminName(name='Alice')],
}
