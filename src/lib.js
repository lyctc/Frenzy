function defaultItemA () {
  return [
    {
      title: 'Navigate with arrow keys',
      labelA: [1],
      childA: [
        {
          title: 'Up/down to move within layers',
          labelA: [1],
          childA: [],
        },
        {
          title: 'Left/right to move between layers',
          labelA: [1],
          childA: [],
        },
        {
          title: 'Enter, then left/right to edit',
          labelA: [1],
          childA: [],
        },
        {
          title: 'Escape to cancel',
          labelA: [1],
          childA: [],
        },
      ],
    },
    {
      title: 'Filter by labels',
      labelA: [2],
      childA: [
        {
          title: 'Shift-1 to add label 1',
          labelA: [2],
          childA: [],
        },
        {
          title: '1 to only show label 1',
          labelA: [2],
          childA: [],
        },
        {
          title: '1 again to show all items',
          labelA: [2],
          childA: [],
        },
      ],
    },
    {
      title: 'Group by labels',
      labelA: [2],
      childA: [
        {
          title: 'G to group by labels',
          labelA: [2],
          childA: [],
        },
        {
          title: 'G again to show structure',
          labelA: [2],
          childA: [],
        },
      ]
    },
    {
      title: 'Plan multiple frenzies',
      labelA: [3],
      childA: [],
    },
    {
      title: 'Get started for free',
      labelA: [3],
      childLabelA: [3],
      childA: [],
    },
  ];
}

function rebalanceItemA(pathItemA, pathDispA, itemA, viewLabelA) {
  // rebalances pathItemA, pathDispA, labelChildA
  let i;
  let r;
  let out = {dispA: [], itemA: [], labelA: []};
  for (i = 0; i < itemA.length; i += 1) {
    let item = JSON.parse(JSON.stringify(itemA[i]));
    let disp = JSON.parse(JSON.stringify(itemA[i]));
    item.pathItemA = pathItemA.concat([i]);
    disp.pathDispA = pathDispA.concat([out.dispA.length]);
    disp.pathItemA = item.pathItemA;

    r = rebalanceItemA(item.pathItemA, disp.pathDispA, item.childA, viewLabelA);

    item.childA = r.itemA;
    item.childLabelA = item.labelA.concat(r.labelA);  // childLabelA has all children's labelA as well as it's own labelA
    out.itemA.push(item);

    disp.childA = r.dispA;
    disp.childLabelA = disp.labelA.concat(r.labelA);  // childLabelA has all children's labelA as well as it's own labelA
    if (viewLabelA.length === 0 || viewLabelA.filter(l => -1 !== disp.childLabelA.indexOf(l)).length > 0) {
      out.dispA.push(disp);
    }

    out.labelA = out.labelA.concat(item.childLabelA);
  }
  return out;
}

function viewLabelHelper(viewLabelA0, label) {
  if (viewLabelA0.indexOf(label) === -1) {
    viewLabelA0.push(label);
  } else {
    viewLabelA0 = viewLabelA0.filter(l => l !== label);
  }
  return viewLabelA0;
}

function labelItemHelper(label, itemA0, pathItemA0, step) {
  // recursively goes to item and toggles label
  if (step === pathItemA0.length - 1) {
    if (itemA0[pathItemA0[step]].labelA.indexOf(label) === -1) {
      itemA0[pathItemA0[step]].labelA.push(label);
    } else {
      itemA0[pathItemA0[step]].labelA = itemA0[pathItemA0[step]].labelA.filter(l => l !== label);
    }
  } else {
    itemA0[pathItemA0[step]].childA = labelItemHelper(label, itemA0[pathItemA0[step]].childA, pathItemA0, step + 1);
  }
  return itemA0;
}

function moveItemHelper(dir, itemA0, pathItemA0, step) {
  let temp;
  if (step === pathItemA0.length - 1) {
    if (dir === 'up' && pathItemA0[step] > 0) {
      temp = itemA0[pathItemA0[step]];
      itemA0.splice(pathItemA0[step], 1);
      itemA0.splice(pathItemA0[step] - 1, 0, temp);
      itemA0[pathItemA0[step]].pathItemA[itemA0[pathItemA0[step]].pathItemA.length - 1] = pathItemA0[step];
      itemA0[pathItemA0[step] - 1].pathItemA[itemA0[pathItemA0[step] - 1].pathItemA.length - 1] = pathItemA0[step] - 1;
    } else if (dir === 'down' && pathItemA0[step] <= itemA0.length - 2) {
      temp = itemA0[pathItemA0[step]];
      itemA0.splice(pathItemA0[step], 1);
      itemA0.splice(pathItemA0[step] + 1, 0, temp);
      itemA0[pathItemA0[step]].pathItemA[itemA0[pathItemA0[step]].pathItemA.length - 1] = pathItemA0[step];
      itemA0[pathItemA0[step] + 1].pathItemA[itemA0[pathItemA0[step] + 1].pathItemA.length - 1] = pathItemA0[step] + 1;
    }
    return itemA0;
  } else {
    itemA0[pathItemA0[step]].childA = moveItemHelper(dir, itemA0[pathItemA0[step]].childA, pathItemA0, step + 1);
  }
  return itemA0;
}

function deleteItemHelper(itemA0, pathItemA0, step) {
  let i;
  if (step === pathItemA0.length - 1) {
    itemA0.splice(pathItemA0[step], 1);
    for (i = 0; i < itemA0.length; i += 1) {
      itemA0[i].pathItemA[step] = i;
    }
    return itemA0;
  } else {
    itemA0[pathItemA0[step]].childA = deleteItemHelper(itemA0[pathItemA0[step]].childA, pathItemA0, step + 1);
  }
  return itemA0;
}

function addItemHelper(title, itemA0, viewLabelA, parentPathItemA, step) {
  // recursively go through itemA and add new item
  // labelA is automatically set to viewLabelA
  if (step === parentPathItemA.length) {
    // parentPathItemA[parentPathItemA.length - 1] doesn't exist yet, so must push
    itemA0.push({
      title: title,
      labelA: viewLabelA,
      childA: [],
    });
    return itemA0;
  } else {
    itemA0[parentPathItemA[step]].childA = addItemHelper(title, itemA0[parentPathItemA[step]].childA, viewLabelA, parentPathItemA, step + 1);
  }
  return itemA0;
}

function updateItemHelper(title, itemA0, pathItemA, step) {
  // recursively go through itemA and update title to new title
  if (step === pathItemA.length - 1) {
    itemA0[pathItemA[step]].title = title;
    return itemA0;
  } else {
    itemA0[pathItemA[step]].childA = updateItemHelper(title, itemA0[pathItemA[step]].childA, pathItemA, step + 1);
  }
  return itemA0;
}

module.exports = {
  defaultItemA: function () {
    return defaultItemA();
  },
  rebalanceItemA: function (pathItemA, pathDispA, itemA0, viewLabelA) {
    return rebalanceItemA(pathItemA, pathDispA, itemA0, viewLabelA);
  },
  viewLabelHelper: function (viewLabelA0, label) {
    return viewLabelHelper(viewLabelA0, label);
  },
  labelItemHelper: function (label, itemA0, pathItemA0, step) {
    return labelItemHelper(label, itemA0, pathItemA0, step);
  },
  moveItemHelper: function (dir, itemA0, pathItemA0, step) {
    return moveItemHelper(dir, itemA0, pathItemA0, step);
  },
  deleteItemHelper: function (itemA0, pathItemA0, step) {
    return deleteItemHelper(itemA0, pathItemA0, step);
  },
  addItemHelper: function (title, itemA0, viewLabelA, pathItemA, step) {
    return addItemHelper(title, itemA0, viewLabelA, pathItemA, step);
  },
  updateItemHelper: function (title, itemA0, pathItemA, step) {
    return updateItemHelper(title, itemA0, pathItemA, step);
  },
};
