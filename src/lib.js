function rebalancePathA(path, itemA0) {
  // rebalances pathA for each item
  let i;
  for (i = 0; i < itemA0.length; i += 1) {
    itemA0[i].pathA = path.concat([i]);
    itemA0[i].childA = rebalancePathA(path.concat([i]), itemA0[i].childA)
  }
  return itemA0;
}

function moveItemHelper(dir, itemA0, posA0, step) {
  let temp;
  if (step === posA0.length - 1) {
    if (dir === 'up' && posA0[step] > 0) {
      temp = itemA0[posA0[step]];
      itemA0.splice(posA0[step], 1);
      itemA0.splice(posA0[step] - 1, 0, temp);
      itemA0[posA0[step]].pathA[itemA0[posA0[step]].pathA.length - 1] = posA0[step];
      itemA0[posA0[step] - 1].pathA[itemA0[posA0[step] - 1].pathA.length - 1] = posA0[step] - 1;
    } else if (dir === 'down' && posA0[step] <= itemA0.length - 2) {
      temp = itemA0[posA0[step]];
      itemA0.splice(posA0[step], 1);
      itemA0.splice(posA0[step] + 1, 0, temp);
      itemA0[posA0[step]].pathA[itemA0[posA0[step]].pathA.length - 1] = posA0[step];
      itemA0[posA0[step] + 1].pathA[itemA0[posA0[step] + 1].pathA.length - 1] = posA0[step] + 1;
    }
    return itemA0;
  } else {
    itemA0[posA0[step]].childA = moveItemHelper(dir, itemA0[posA0[step]].childA, posA0, step + 1);
  }
  return itemA0;
}

function deleteItemHelper(itemA0, posA0, step) {
  let i;
  if (step === posA0.length - 1) {
    itemA0.splice(posA0[step], 1);
    for (i = 0; i < itemA0.length; i += 1) {
      itemA0[i].pathA[step] = i;
    }
    return itemA0;
  } else {
    itemA0[posA0[step]].childA = deleteItemHelper(itemA0[posA0[step]].childA, posA0, step + 1);
  }
  return itemA0;
}

function addItemHelper(title, itemA0, posA, step) {
  // recursively go through itemA and add new item
  if (step === posA.length - 1) {
    // posA[posA.length - 1] doesn't exist yet, so must push
    itemA0.push({
      pathA: posA,
      iid: 0,
      title: title,
      labelA: [],
      childA: [],
    });
    return itemA0;
  } else {
    itemA0[posA[step]].childA = addItemHelper(title, itemA0[posA[step]].childA, posA, step + 1);
  }
  return itemA0;
}

function updateItemHelper(title, itemA0, posA, step) {
  // recursively go through itemA and update title to new title
  if (step === posA.length - 1) {
    itemA0[posA[step]].title = title;
    return itemA0;
  } else {
    itemA0[posA[step]].childA = updateItemHelper(title, itemA0[posA[step]].childA, posA, step + 1);
  }
  return itemA0;
}

module.exports = {
  rebalancePathA: function (path, itemA0) {
    return rebalancePathA(path, itemA0);
  },
  moveItemHelper: function (dir, itemA0, posA0, step) {
    return moveItemHelper(dir, itemA0, posA0, step);
  },
  deleteItemHelper: function (itemA0, posA0, step) {
    return deleteItemHelper(itemA0, posA0, step);
  },
  addItemHelper: function (title, itemA0, posA, step) {
    return addItemHelper(title, itemA0, posA, step);
  },
  updateItemHelper: function (title, itemA0, posA, step) {
    return updateItemHelper(title, itemA0, posA, step);
  },
};
