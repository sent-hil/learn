items  = []
doc = {
  :id     => 1,
  :script => "ctx._source.count += 1",
  :upsert => {
    'count'      => 1,
    '@timestamp' => Time.now.utc,
  }
}

items << doc

index = Tire.index 'whatever'
index.bulk :update, vms
