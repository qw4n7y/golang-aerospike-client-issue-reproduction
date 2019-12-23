local function does_list_contain(l, item)
  for value in list.iterator(l) do
    if value == item then return true end
  end
  return false
end

local function rec_to_map(rec)
  local xrec = map()
  for i, bin_name in ipairs(record.bin_names(rec)) do
    xrec[bin_name] = rec[bin_name]
  end
  return xrec
end

function myUDF(stream, myList)
  info("myUDF = %s (%s)", myList, type(myList))
  
  local function my_filter(rec)
    return does_list_contain(myList, rec['foo'])
  end

  return stream:filter(my_filter):map(rec_to_map)
end