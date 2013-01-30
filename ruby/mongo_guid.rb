# Problem: Unique document id for objects, without using ugly uids.
#
# Solution: Have a seperate 'counters' collection, use findAndModify to find
#   the counters id and increment by random number it at same time.
#
#   Save object record with returned counter+random number as id.

require 'mongo'

def conn
  Mongo::Connection.new
end

def db
  conn['guid']
end

def coll
  db['counters']
end

# Insert counter at start.
def insert_counter
  coll.insert({'name' => 'guid', 'counter'=> 0})
end

def documents
  coll.find('name' => 'guid').to_a
end

def find_and_modify_counter(random_number=rand(100))
  coll.find_and_modify(:query => {'name' => 'guid'},
                       :update => {'$inc' => {'counter' => random_number}})
end
