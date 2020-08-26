package sectorstorage

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/specs-actors/actors/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
)

type xgP1P2Selector struct {
	index      stores.SectorIndex
	sector     abi.SectorID
	alloc      stores.SectorFileType
	allowFetch bool
	hostName   string
}

func newXgP1P2Selector(index stores.SectorIndex, sector abi.SectorID, alloc stores.SectorFileType, allowFetch bool, hostName string) *xgP1P2Selector {
	return &xgP1P2Selector{
		index:      index,
		sector:     sector,
		alloc:      alloc,
		allowFetch: allowFetch,
		hostName:   hostName,
	}
}

func (s *xgP1P2Selector) Ok(ctx context.Context, task sealtasks.TaskType, spt abi.RegisteredSealProof, whnd *workerHandle) (bool, error) {
	tasks, err := whnd.w.TaskTypes(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}
	if _, supported := tasks[task]; !supported {
		return false, nil
	}

	log.Infof("-------xg------ whnd: %s, selector: %s", whnd.info.Hostname, s.hostName)

	if whnd.info.Hostname == s.hostName {
		return true, nil
	}

	//paths, err := whnd.w.Paths(ctx)
	//if err != nil {
	//	return false, xerrors.Errorf("getting worker paths: %w", err)
	//}
	//
	//have := map[stores.ID]struct{}{}
	//for _, path := range paths {
	//	have[path.ID] = struct{}{}
	//}
	//
	//best, err := s.index.StorageFindSector(ctx, s.sector, s.alloc, spt, s.allowFetch)
	//if err != nil {
	//	return false, xerrors.Errorf("finding best storage: %w", err)
	//}
	//
	//for _, info := range best {
	//	if _, ok := have[info.ID]; ok {
	//		return true, nil
	//	}
	//}

	return false, nil
}

func (s *xgP1P2Selector) Cmp(ctx context.Context, task sealtasks.TaskType, a, b *workerHandle) (bool, error) {
	return a.active.utilization(a.info.Resources) < b.active.utilization(b.info.Resources), nil
}

var _ WorkerSelector = &xgP1P2Selector{}
